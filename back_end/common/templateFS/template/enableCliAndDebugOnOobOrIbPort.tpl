####enable debug port on OOB/IB port####
<rpc message-id="1" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
  <edit-config>
    <target>
      <running/>
    </target>
    <config>
      <system xmlns="urn:ietf:params:xml:ns:yang:ietf-system">
        <management xmlns="http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-ietf-system-aug">
          <debug>
            <ip_itf ns0:operation="merge" xmlns:ns0="urn:ietf:params:xml:ns:netconf:base:1.0">
              <enable>true</enable>
            </ip_itf>
          </debug>
        </management>
      </system>
    </config>
  </edit-config>
</rpc>



####enable CLI port on OOB/IB port####
<rpc message-id="1" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
  <edit-config>
    <target>
      <running/>
    </target>
    <config xmlns="http://tail-f.com/ns/config/1.0">
      <system xmlns="urn:ietf:params:xml:ns:yang:ietf-system">
        <management xmlns="http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-ietf-system-aug">
          <cli>
            <transport>
              <ssh>
                <ip_itf>
                  <enable>true</enable>
                </ip_itf>
              </ssh>
            </transport>
          </cli>
        </management>
      </system>
    </config>
  </edit-config>
</rpc>