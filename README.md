# euro-rates

Get Euro rates from XML and parse it with Golang

---

**USAGE**

`install golang`

`git clone https://github.com/localizepic/euro-rates.git`

`go run euro-rates/src/get-euro-rates.go`

---

**SOURCE URL** 
`https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml?timestamp`

---

**XML structure**


`<gesmes:Envelope xmlns:gesmes="http://www.gesmes.org/xml/2002-08-01" xmlns="http://www.ecb.int/vocabulary/2002-08-01/eurofxref">
        
        <gesmes:subject>Reference rates</gesmes:subject>
        <gesmes:Sender>
                <gesmes:name>European Central Bank</gesmes:name>
        </gesmes:Sender>
        <Cube>
                <Cube time='2019-03-08'>
                        <Cube currency='USD' rate='1.1222'/>
                        <Cube currency='JPY' rate='124.72'/>
                        <Cube currency='GBP' rate='0.85905'/>
                        <Cube currency='HUF' rate='315.95'/>
                        <Cube currency='PLN' rate='4.3002'/>
                        <Cube currency='RON' rate='4.7448'/>
                </Cube>
        </Cube>
</gesmes:Envelope>`

---

**RESULT example**

`Country Code: USD        Rate: 1.1222`

`Country Code: JPY        Rate: 124.72`

`Country Code: BGN        Rate: 1.9558`

