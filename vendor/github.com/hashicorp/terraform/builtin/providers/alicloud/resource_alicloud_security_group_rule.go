package alicloud

import (
	"fmt"
	"log"
	"strings"

	"github.com/denverdino/aliyungo/ecs"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAliyunSecurityGroupRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliyunSecurityGroupRuleCreate,
		Read:   resourceAliyunSecurityGroupRuleRead,
		Delete: resourceAliyunSecurityGroupRuleDelete,

		Schema: map[string]*schema.Schema{
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateSecurityRuleType,
				Description:  "Type of rule, ingress (inbound) or egress (outbound).",
			},

			"ip_protocol": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateSecurityRuleIpProtocol,
			},

			"nic_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validateSecurityRuleNicType,
			},

			"policy": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validateSecurityRulePolicy,
			},

			"port_range": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"priority": {
				Type:         schema.TypeInt,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validateSecurityPriority,
			},

			"security_group_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"cidr_ip": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "0.0.0.0/0",
			},

			"source_security_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"source_group_owner_account": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliyunSecurityGroupRuleCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AliyunClient).ecsconn

	ruleType := d.Get("type").(string)
	sgId := d.Get("security_group_id").(string)
	ptl := d.Get("ip_protocol").(string)
	port := d.Get("port_range").(string)

	var autherr error
	switch GroupRuleDirection(ruleType) {
	case GroupRuleIngress:
		args, err := buildAliyunSecurityIngressArgs(d, meta)
		if err != nil {
			return err
		}
		autherr = conn.AuthorizeSecurityGroup(args)
	case GroupRuleEgress:
		args, err := buildAliyunSecurityEgressArgs(d, meta)
		if err != nil {
			return err
		}
		autherr = conn.AuthorizeSecurityGroupEgress(args)
	default:
		return fmt.Errorf("Security Group Rule must be type 'ingress' or type 'egress'")
	}

	if autherr != nil {
		return fmt.Errorf(
			"Error authorizing security group rule type %s: %s",
			ruleType, autherr)
	}

	d.SetId(sgId + ":" + ruleType + ":" + ptl + ":" + port)
	return resourceAliyunSecurityGroupRuleRead(d, meta)
}

func resourceAliyunSecurityGroupRuleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AliyunClient)
	parts := strings.Split(d.Id(), ":")
	sgId := parts[0]
	types := parts[1]
	ip_protocol := parts[2]
	port_range := parts[3]
	rule, err := client.DescribeSecurityGroupRule(sgId, types, ip_protocol, port_range)

	if err != nil {
		if notFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error SecurityGroup rule: %#v", err)
	}
	log.Printf("[WARN]sg %s, type %s, protocol %s, port %s, rule %#v", sgId, types, ip_protocol, port_range, rule)
	d.Set("type", rule.Direction)
	d.Set("ip_protocol", strings.ToLower(string(rule.IpProtocol)))
	d.Set("nic_type", rule.NicType)
	d.Set("policy", strings.ToLower(string(rule.Policy)))
	d.Set("port_range", rule.PortRange)
	d.Set("priority", rule.Priority)
	d.Set("security_group_id", sgId)
	//support source and desc by type
	if GroupRuleDirection(types) == GroupRuleIngress {
		d.Set("cidr_ip", rule.SourceCidrIp)
		d.Set("source_security_group_id", rule.SourceGroupId)
		d.Set("source_group_owner_account", rule.SourceGroupOwnerAccount)
	} else {
		d.Set("cidr_ip", rule.DestCidrIp)
		d.Set("source_security_group_id", rule.DestGroupId)
		d.Set("source_group_owner_account", rule.DestGroupOwnerAccount)
	}

	return nil
}

func resourceAliyunSecurityGroupRuleDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AliyunClient)
	args, err := buildAliyunSecurityIngressArgs(d, meta)

	if err != nil {
		return err
	}
	revokeArgs := &ecs.RevokeSecurityGroupArgs{
		AuthorizeSecurityGroupArgs: *args,
	}
	return client.RevokeSecurityGroup(revokeArgs)
}

func buildAliyunSecurityIngressArgs(d *schema.ResourceData, meta interface{}) (*ecs.AuthorizeSecurityGroupArgs, error) {
	conn := meta.(*AliyunClient).ecsconn

	args := &ecs.AuthorizeSecurityGroupArgs{
		RegionId: getRegion(d, meta),
	}

	if v := d.Get("ip_protocol").(string); v != "" {
		args.IpProtocol = ecs.IpProtocol(v)
	}

	if v := d.Get("port_range").(string); v != "" {
		args.PortRange = v
	}

	if v := d.Get("policy").(string); v != "" {
		args.Policy = ecs.PermissionPolicy(v)
	}

	if v := d.Get("priority").(int); v != 0 {
		args.Priority = v
	}

	if v := d.Get("nic_type").(string); v != "" {
		args.NicType = ecs.NicType(v)
	}

	if v := d.Get("cidr_ip").(string); v != "" {
		args.SourceCidrIp = v
	}

	if v := d.Get("source_security_group_id").(string); v != "" {
		args.SourceGroupId = v
	}

	if v := d.Get("source_group_owner_account").(string); v != "" {
		args.SourceGroupOwnerAccount = v
	}

	sgId := d.Get("security_group_id").(string)

	sgArgs := &ecs.DescribeSecurityGroupAttributeArgs{
		SecurityGroupId: sgId,
		RegionId:        getRegion(d, meta),
	}

	_, err := conn.DescribeSecurityGroupAttribute(sgArgs)
	if err != nil {
		return nil, fmt.Errorf("Error get security group %s error: %#v", sgId, err)
	}

	args.SecurityGroupId = sgId

	return args, nil
}

func buildAliyunSecurityEgressArgs(d *schema.ResourceData, meta interface{}) (*ecs.AuthorizeSecurityGroupEgressArgs, error) {
	conn := meta.(*AliyunClient).ecsconn

	args := &ecs.AuthorizeSecurityGroupEgressArgs{
		RegionId: getRegion(d, meta),
	}

	if v := d.Get("ip_protocol").(string); v != "" {
		args.IpProtocol = ecs.IpProtocol(v)
	}

	if v := d.Get("port_range").(string); v != "" {
		args.PortRange = v
	}

	if v := d.Get("policy").(string); v != "" {
		args.Policy = ecs.PermissionPolicy(v)
	}

	if v := d.Get("priority").(int); v != 0 {
		args.Priority = v
	}

	if v := d.Get("nic_type").(string); v != "" {
		args.NicType = ecs.NicType(v)
	}

	if v := d.Get("cidr_ip").(string); v != "" {
		args.DestCidrIp = v
	}

	if v := d.Get("source_security_group_id").(string); v != "" {
		args.DestGroupId = v
	}

	if v := d.Get("source_group_owner_account").(string); v != "" {
		args.DestGroupOwnerAccount = v
	}

	sgId := d.Get("security_group_id").(string)

	sgArgs := &ecs.DescribeSecurityGroupAttributeArgs{
		SecurityGroupId: sgId,
		RegionId:        getRegion(d, meta),
	}

	_, err := conn.DescribeSecurityGroupAttribute(sgArgs)
	if err != nil {
		return nil, fmt.Errorf("Error get security group %s error: %#v", sgId, err)
	}

	args.SecurityGroupId = sgId

	return args, nil
}
