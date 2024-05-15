package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "net"
    "time"
)

const hl7Message = `MSH|^~\&|SendingApplication|SendingFacility|ReceivingApplication|ReceivingFacility|20240515094500||ADT^A01|MSG00001|P|2.5
EVN|A01|20240515094500
PID|1||123456^^^Hospital^MR||Doe^John^^^^^L||19800101|M|||123 Fake St^^Faketown^CA^12345^USA||123-456-7890|||S||123456789|987-65-4321
PV1|1|I|2000^2012^01||||004777^Smith^John^A|||SUR||||||||ADM|A0|`

type Response struct {
    Request  string `json:"request"`
    Response string `json:"response"`
}

func main() {
    ip := flag.String("ip", "127.0.0.1", "Target IP address")
    port := flag.String("port", "2575", "Target port")
    flag.Parse()

    target := fmt.Sprintf("%s:%s", *ip, *port)

    conn, err := net.Dial("tcp", target)
    if err != nil {
        fmt.Printf("Failed to connect: %v\n", err)
        return
    }
    defer conn.Close()

    fmt.Println("Connected to FHIR server")

    _, err = conn.Write([]byte(hl7Message))
    if err != nil {
        fmt.Printf("Failed to send message: %v\n", err)
        return
    }

    fmt.Println("HL7 message sent")
    conn.SetReadDeadline(time.Now().Add(5 * time.Second))

    buffer := make([]byte, 4096)
    n, err := conn.Read(buffer)
    if err != nil {
        fmt.Printf("Failed to receive response: %v\n", err)
        return
    }

    response := Response{
        Request:  hl7Message,
        Response: string(buffer[:n]),
    }

    responseJSON, err := json.MarshalIndent(response, "", "  ")
    if err != nil {
        fmt.Printf("Failed to marshal response: %v\n", err)
        return
    }

    fmt.Println(string(responseJSON))
}
