---
title: Supply-chain attack on security software
date: "2021-04-25T14:40:24+02:00"
tags:
- security
- news
- 100daystooffload
incoming:
- url: https://chaos.social/@zerok/106126026592321029
---

Last Friday [the news hit](https://arstechnica.com/gadgets/2021/04/hackers-backdoor-corporate-password-manager-and-steal-customer-data/) that Passwordstate, a password manager that is mostly used in corporate environments for shared password lists, has been successfully attacked through a their update mechanism. Around 29,000 users may have been affected here with potentially all secrets stored in their respective Passwordstate installation having been compromised. 

Details on the attack itself are sparse but it looks like the attackers were able to manipulate an update file provided by [Click Studios](https://www.clickstudios.com.au/), the company behind Passwordstate. An administrator that downloaded that update file from the official update site between April 20 and April 22 and install the update file hasn’t only installed the update but also a malware client on their servers that submits various data sets to the attackers. 

According to an [advisory by Click Studios](https://www.clickstudios.com.au/advisories/Incident_Management_Advisory-01-20210424.pdf) affects machines have posted following data back to the attackers:

> Analysis of compromised data indicates the following information is posted back:
> Computer Name, User Name, Domain Name, Current Process Name, Current Process Id, All running Processes name and ID, All running services name, display name and status, Passwordstate instance’s Proxy Server Address, Username and Password
> 
> The following fields in Passwordstate instance’s password table is posted back:
> Title, UserName, Description, GenericField1, GenericField2, GenericField3, Notes, URL, Password
> 
> The Domain Name and Host name aren’t extracted as part of this. Although the encryption key and database connection string are used to process data via hooking into the Passwordstate Service process, there is no evidence of encryption keys or database connection strings being posted to the bad actor CDN network.

I’m really curious how the update mechanism of Passwordstate workers here. Are checksums require? Are these checksums signed? With what key? Does the update mechanism within Passwordstate enforce the validation of such checksums and their signatures? Can administrators skips checks during an update here? Or even worse: Is all of that optional and mostly left to administrators themselves?

Click Studios is far from the only supplier that has recently been compromised through their update mechanism. Sadly, the whole deployment and update process seems still be treated like something that everyone would love to just get over with all the time. It’s just really sad to also see successful attacks against these processes also in security relevant software. Password manages are immensely useful and important, but if they are compromised they allow attackers not only access to a single but a multitude of environments.

Can we as an industry and community take software updates and update mechanisms a bit more serious?!
