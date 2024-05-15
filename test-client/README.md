## Test-client 

Simple golang application for testing HL7 FHIR servers to confirm honeypot function

## Example output
```
{"level":"info","message":"Connection found","timestamp":"15 May 24 17:47 UTC","src_port":"58642","src_ip":"127.0.0.1","data":"TVNIfF5+XCZ8U2VuZGluZ0FwcGxpY2F0aW9ufFNlbmRpbmdGYWNpbGl0eXxSZWNlaXZpbmdBcHBsaWNhdGlvbnxSZWNlaXZpbmdGYWNpbGl0eXwyMDI0MDUxNTA5NDUwMHx8QURUXkEwMXxNU0cwMDAwMXxQfDIuNQpFVk58QTAxfDIwMjQwNTE1MDk0NTAwClBJRHwxfHwxMjM0NTZeXl5Ib3NwaXRhbF5NUnx8RG9lXkpvaG5eXl5eXkx8fDE5ODAwMTAxfE18fHwxMjMgRmFrZSBTdF5eRmFrZXRvd25eQ0FeMTIzNDVeVVNBfHwxMjMtNDU2LTc4OTB8fHxTfHwxMjM0NTY3ODl8OTg3LTY1LTQzMjEKUFYxfDF8SXwyMDAwXjIwMTJeMDF8fHx8MDA0Nzc3XlNtaXRoXkpvaG5eQXx8fFNVUnx8fHx8fHx8QURNfEEwfA=="}
{
  "request": "MSH|^~\\\u0026|SendingApplication|SendingFacility|ReceivingApplication|ReceivingFacility|20240515094500||ADT^A01|MSG00001|P|2.5\nEVN|A01|20240515094500\nPID|1||123456^^^Hospital^MR||Doe^John^^^^^L||19800101|M|||123 Fake St^^Faketown^CA^12345^USA||123-456-7890|||S||123456789|987-65-4321\nPV1|1|I|2000^2012^01||||004777^Smith^John^A|||SUR||||||||ADM|A0|",
  "response": "MSH|^~\\\u0026|SENDING_APPLICATION|SENDING_FACILITY|RECEIVING_APPLICATION|RECEIVING_FACILITY|20110614075841||ACK|1407511|P|2.3||||\n        MSA|AA|1407511|Success||"
}```
