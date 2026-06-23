  <action xmlns="urn:ietf:params:xml:ns:yang:1">
    <hardware-state xmlns="urn:ietf:params:xml:ns:yang:ietf-hardware">
      <component>
        <name>Chassis</name>
        <software xmlns="urn:bbf:yang:bbf-software-image-management-one-dot-one">
          <software>
            <name>application_software</name>
            <revisions>
              <revision>
                <name>{{.SoftwareName}}</name>
                <download xmlns="http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-software-image-management-extension">
                  <config-download>
                    <url>{{.SoftwareUrl}}</url>
                  </config-download>
                </download>
              </revision>
            </revisions>
          </software>
        </software>
      </component>
    </hardware-state>
  </action>