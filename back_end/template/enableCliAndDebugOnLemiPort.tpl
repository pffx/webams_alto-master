####enable debug port on LEMI port####
<rpc message-id="1" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
  <edit-config>
    <target>
      <running/>
    </target>
    <config>
      <system xmlns="urn:ietf:params:xml:ns:yang:ietf-system">
        <management xmlns="http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-ietf-system-aug">
          <debug>
            <lemi>
              <enable>true</enable>
            </lemi>
          </debug>
        </management>
      </system>
    </config>
  </edit-config>
</rpc>



####enable CLI port on LEMI port####
<rpc message-id="1" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
  <edit-config>
    <target>
      <running/>
    </target>
    <config>
      <system xmlns="urn:ietf:params:xml:ns:yang:ietf-system">
        <management xmlns="http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-ietf-system-aug">
          <cli>
            <transport>
              <ssh>
                <lemi>
                  <enable>true</enable>
                </lemi>
              </ssh>
            </transport>
          </cli>
        </management>
      </system>
    </config>
  </edit-config>
</rpc>