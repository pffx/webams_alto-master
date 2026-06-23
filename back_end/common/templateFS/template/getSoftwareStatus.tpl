<get xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
  <filter type="subtree">
    <hardware-state xmlns="urn:ietf:params:xml:ns:yang:ietf-hardware">
      <component>
        <name>Chassis</name>
        <software xmlns="urn:bbf:yang:bbf-software-image-management-one-dot-one">
          <software>
            <name>application_software</name>
          </software>
        </software>
      </component>
    </hardware-state>
  </filter>
</get>