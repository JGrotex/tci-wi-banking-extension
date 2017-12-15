# TCI WI Banking Extension
first Version with just IBAN validation, more to come ...

Here you can test the running Service on TIBCO Cloud Integration & Web Integrator and including this Extension using Swagger UI
http://www.godev.de/GOLib/swaggUI.html

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
