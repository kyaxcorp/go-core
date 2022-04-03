package fcm

/*
Sa fi instanta autonoma si separata
sa lucreze complete izolat... ca un poroces separat?!
sa lucrez in acelasi proces la un component existent
sa funcitoneze cu cockroach si comunicarea crearea notificarilor sa fie realizata prin cockroach?!
Trigger-ul trebuie oricum sa fie facut spre un proces existent...!?
Corect a r fi ca procesul sa primeasca request-ul... chiar si prin eveniment daca este necesar!
Si acesta deja sa adauge in DB informatia despre notificare


In general acesta trbeuie sa primeasca deja continutul final + token-ul dispozitivului!
Decizia "cui trebuie sa fie transmis mesajul" trbeuie sa faca acela care vrea sa trnasmita....



- Tabel si model in DB
- Serviciul care va citi din Tabel si va transmite mesajele
- Serviciul dat trebuie cumvai sa fie notificat ca are new messages to send! ca sa nu scaneze tot tabelul integral poate
acesta sa fie notificat de unde sa transmita?

De asemenea cind transmite, poate sa facem lock la proces mesajul care este in curs de transmitere... acesta de asemenea
va avea lock_ttl vreo 300 secunde sau ceva de genul..., caci in cazul in care nu se este transmis... un alt proces deja
sa poata sa-l preia pe acel vechi!


We can send 2 types of messages:
- Notification Messages FCM SDK AUTO!
- Data messages, handled by the APP

https://firebase.google.com/docs/cloud-messaging/concept-options?authuser=1



*/
