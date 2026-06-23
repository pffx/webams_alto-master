<action xmlns="urn:ietf:params:xml:ns:yang:1">
  <hardware-state xmlns="urn:ietf:params:xml:ns:yang:ietf-hardware">
    <component>
      <name>Chassis</name>
      <software xmlns="urn:bbf:yang:bbf-software-image-management-one-dot-one">
        <software>
          <name>application_software</name>
          <download>
            <download-software>
              <url>{{.SoftwareUrl}}</url>
              <name>{{.SoftwareName}}</name>
            </download-software>
          </download>
        </software>
      </software>
    </component>
  </hardware-state>
</action>