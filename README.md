# medpot
Is a honeypot that tries to emulate HL7 / FHIR honeypot

* Listening on port 2575

## Installation

1. You need go
2. Install dep
3. Run dep ensure


## Log
Logfiles are  located at `/var/log/medpot.log`

<b>Example</b>
```
{
    "level":"info",
    "message":"Connection found",
    "time":"2018.09.09 17:20:49",
    "port":"57905",
    "ip":"127.0.0.1",
    "data":"TVMpEUk98S0FURV5TTUlUSF5FEK"
}
```


