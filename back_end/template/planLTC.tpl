  <edit-config>
    <target>
      <running/>
    </target>
    <config>
      <hardware xmlns="urn:ietf:params:xml:ns:yang:ietf-hardware">
        <component xmlns:ns0="urn:ietf:params:xml:ns:netconf:base:1.0" ns0:operation="create">
          <name>lt-{{.LtIndex}}</name>
          <parent>Slot-Lt-{{.LtIndex}}</parent>
          <parent-rel-pos>1</parent-rel-pos>
          <mfg-name>ALCL</mfg-name>
          <model-name xmlns="urn:bbf:yang:bbf-hardware-extension">{{.ModelName}}</model-name>
          <class xmlns:nokia-hwi="http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-hardware-identities">nokia-hwi:lt</class>
        </component>
      </hardware>
    </config>
  </edit-config>