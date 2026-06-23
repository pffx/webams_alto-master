<?xml version="1.0" ?>
<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">
    <xsl:strip-space elements="*"/>
    <xsl:output method="xml" indent="yes"/>

    <!-- default rule -->
    <xsl:template match="*">
        <xsl:copy>
            <xsl:copy-of select="@*"/>
            <xsl:apply-templates/>
        </xsl:copy>
    </xsl:template>

    <xsl:template match="*[name() = 'capabilities']">
        <xsl:choose>
            <xsl:when
                    test="namespace-uri() = 'http://www.nokia.com/ns/dpu/nokia-capability' and parent::*[name() = 'config']">
            </xsl:when>
            <xsl:otherwise>
                <xsl:copy>
                    <xsl:apply-templates select="@* | node()"/>
                </xsl:copy>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>

    <xsl:template match="*[name() = 'login-user']">
        <xsl:choose>
            <xsl:when
                    test="namespace-uri() = 'http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-aaa-hidden' and parent::*[name() = 'config'] ">
            </xsl:when>
            <xsl:otherwise>
                <xsl:copy>
                    <xsl:apply-templates select="@* | node()"/>
                </xsl:copy>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>

    <xsl:template match="*[name() = 'object-index']">
        <xsl:choose>
            <xsl:when test=" 'y' = 'y' ">
            </xsl:when>
            <xsl:otherwise>
                <xsl:copy>
                    <xsl:apply-templates select="@* | node()"/>
                </xsl:copy>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>
    <xsl:template match="*[name() = 'aaa']">
        <xsl:choose>
            <xsl:when test="namespace-uri() = 'http://tail-f.com/ns/aaa/1.1' and parent::*[name() = 'config'] ">
            </xsl:when>
            <xsl:otherwise>
                <xsl:copy>
                    <xsl:apply-templates select="@* | node()"/>
                </xsl:copy>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>
    <xsl:template match="*[name() = 'dsc-server']">
        <xsl:choose>
            <xsl:when
                    test="namespace-uri() = 'urn:ietf:params:xml:ns:yang:eodb-vnf-config' and parent::*[name() = 'config'] ">
            </xsl:when>
            <xsl:otherwise>
                <xsl:copy>
                    <xsl:apply-templates select="@* | node()"/>
                </xsl:copy>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>
    <xsl:template match="*[name() = 'config'
    and namespace-uri() = 'http://tail-f.com/ns/config/1.0'
]">
        <xsl:element name="config" namespace="urn:ietf:params:xml:ns:netconf:base:1.0">
            <xsl:copy-of select="child::*"/>
        </xsl:element>
    </xsl:template>
</xsl:stylesheet>