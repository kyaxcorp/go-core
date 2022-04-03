1. sa fac log-uri la websocket client
2. sa verific writer-ul
5. sa vedem unde setam evenimentul onError... pentru ca logic el ar trebui sa fie pentru toate erorile din client..
6. 
Sa setez event-urile

onError
onSendError
onSend
onConnect

7.
Statistics?!...
cind s-a transmis in bytes
cind am primit
cind text a fost transmis
cind binary a fost transmis
cit json
cite mesaje au fost primite
cite mesaje au fost transmise
in cit timp?!
8.
sa fac o functie cu LATENCY... in care aflam in cit timp raspunsul vine de la server?!..
9 Sa fac before send si after send... in functiile de transmitere... evenimente comode
10. sa fac mecanism de cunoastere daca mesajul x a fost transmis cu success sau nu.. trebuie sa fie raspuns din partea la server!
11. TLS Security with custom certificates
12. Accept self signed certs
    - maybe somehow identify these self signed... even if they are tampered...?! maybe by checksum or something!...
13.  De facut functiile de citire a statisticii
14.  Add other statistics like: nr of disconnections, reconnections