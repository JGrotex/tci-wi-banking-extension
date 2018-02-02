# TCI WI Banking Extension

[![Go Report Card](https://goreportcard.com/badge/github.com/JGrotex/tci-wi-banking-extension)](https://goreportcard.com/report/github.com/JGrotex/tci-wi-banking-extension) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

first Version with just IBAN validation, more to come ...

Usually you have to pay for those kind of services, with this Extension you can first of all host your own Version, and apply your own addantional Validations as well.

Here you can test the running Service on TIBCO Cloud Integration & Web Integrator and including this Extension using Swagger UI
http://www.godev.de/GOLib/swaggUI.html

Will add the full TIBCO Cloud Integration App here soon, as the WebUI comes now with a APP Export Functionality.

Attached ZIP contains the first release v.1.2 and can just uploaded under 
TIBCO Cloud Integration - Extensions

## Activities
available Activities so far
### IBAN Validation
Validation of an IBAN String against IBAN Country specific Patterns and Checksum verification.  

#### Graphical Webintegrator Flow in TCI
![TCI Webintegrator flow image](screenshots/Banking-IBAN.png?raw=true "TCI WI Banking IBAN validation Screenshot")

#### Interface
Input
- IBAN (String)

Output
- result (String) "valid" / "invalid"

<hr>
<sub><b>Note:</b> more TCI Extensions can be found here: https://tibcosoftware.github.io/tci-awesome/ </sub>
