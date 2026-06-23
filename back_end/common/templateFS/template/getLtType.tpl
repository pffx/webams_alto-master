<get xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
  <filter type="subtree">
    <hardware-state xmlns="urn:ietf:params:xml:ns:yang:ietf-hardware">
      <component>
        <name>lt-{{.LtIndex}}</name>
        <model-name></model-name>
      </component>
    </hardware-state>
  </filter>
</get>