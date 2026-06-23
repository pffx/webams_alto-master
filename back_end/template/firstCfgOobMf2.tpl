####configure IP####
<rpc message-id="1" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
  <edit-config>
    <target>
      <running/>
    </target>
    <config>
      <interfaces xmlns="urn:ietf:params:xml:ns:yang:ietf-interfaces">
        <interface xmlns:ns0="urn:ietf:params:xml:ns:netconf:base:1.0" ns0:operation="delete">
          <name>outband</name>
		  <type xmlns:ianaift="urn:ietf:params:xml:ns:yang:iana-if-type">ianaift:ipForward</type>
        </interface>
      </interfaces>
	  <interfaces xmlns="urn:ietf:params:xml:ns:yang:ietf-interfaces">
        <interface xmlns:ns0="urn:ietf:params:xml:ns:netconf:base:1.0" ns0:operation="delete">
          <name>inband</name>
		  <type xmlns:ianaift="urn:ietf:params:xml:ns:yang:iana-if-type">ianaift:ipForward</type>
        </interface>
      </interfaces>
	  <interfaces xmlns="urn:ietf:params:xml:ns:yang:ietf-interfaces">
        <interface xmlns:ns0="urn:ietf:params:xml:ns:netconf:base:1.0" ns0:operation="delete">
          <name>sinband</name>
		  <type xmlns:bbfift="urn:bbf:yang:bbf-if-type">bbfift:vlan-sub-interface</type>
        </interface>
      </interfaces>
      <interfaces xmlns="urn:ietf:params:xml:ns:yang:ietf-interfaces">
        <interface>
          <name>outband</name>
          <type xmlns:ianaift="urn:ietf:params:xml:ns:yang:iana-if-type">ianaift:ipForward</type>
          <ipv4 xmlns="urn:ietf:params:xml:ns:yang:ietf-ip">
            <address>
              <ip>{{.IpAddr}}</ip>
              <netmask>{{.NetMask}}</netmask>
            </address>
          </ipv4>
        </interface>
      </interfaces>
    </config>
  </edit-config>
</rpc>


####configure route####
<rpc message-id="1" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
  <edit-config>
    <target>
      <running/>
    </target>
    <config>
      <routing xmlns="urn:ietf:params:xml:ns:yang:ietf-routing">
        <control-plane-protocols>
          <control-plane-protocol xmlns:ns0="urn:ietf:params:xml:ns:netconf:base:1.0" ns0:operation="delete">
            <type>static</type>
            <name>outband</name>
          </control-plane-protocol>
        </control-plane-protocols>
      </routing>
      <routing xmlns="urn:ietf:params:xml:ns:yang:ietf-routing">
        <control-plane-protocols>
          <control-plane-protocol>
            <type>static</type>
            <name>outband</name>
            <static-routes>
              <ipv4 xmlns="urn:ietf:params:xml:ns:yang:ietf-ipv4-unicast-routing">
                <route>
                  <destination-prefix>0.0.0.0/0</destination-prefix>
                  <next-hop>
                    <next-hop-address>{{.DefaultRoute}}</next-hop-address>
                  </next-hop>
                </route>
              </ipv4>
            </static-routes>
          </control-plane-protocol>
        </control-plane-protocols>
      </routing>
    </config>
  </edit-config>
</rpc>


