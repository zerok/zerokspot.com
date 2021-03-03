---
date: '2005-06-23T12:00:00-00:00'
language: en
tags:
- java
title: RMI Tutorial
---


# RMI Tutorial

Was ist RMI? RMI bedeuted "Remote Method Invokation" und dient dazu, Dienste auf einem Rechner anzubieten, die dann von einem anderen Rechner in Anspruch genommen werden können. Bei diesen Diensten handelt es sich um Java-Objekte die dann am Clientrechner direkt in lokale Variablen eingebunden werden können.

Wenn jedoch eine Methode dieses Objektes ausgeführt werden soll, dann wird diese auf dem "Server" ausgeführt und das Ergebnis dann an den Client übermittelt.



-------------------------------



## Die Server-Seite

### Interface

Auf der Serverseite sollte zuerst einmal ein Interface angelegt werden, damit der Client genau weiß, welche Methoden und Variablen ihm von dem am Server angebotenen Dienst zur Verfügung stehen. In diesem Beispiel bieten wir einfach eine Methode <span class="method">getValue()</span> an, die einen Integer an den Client zurückgeben soll.

Folglich würde das Interface unseres Servers dann in etwa so aussehen:

<pre class="code">
import java.rmi.*;
public interface Server extends Remote
{
	int getValue() throws RemoteException;
}
</pre>

Wie man hier sieht, erweitert unser <span class="interface">Server</span>-Interface das <span class="interface">java.rmi.Remote</span>-Interface. Dadurch wird klargestellt, dass alle Objekte, die dieses Interface implementieren (bzw. ein Interface implementieren, das dieses erweitert) Objekte sind, die von einer nicht-lokalen VM verwendet werden können. Und da bei einem Transfer immer etwas schief gehen kann, müssen alle Methoden (sowie auch der Konstruktor) eine java.rmi.RemoteException werfen können.

<div class="note">
	
Unsere Methode <span class="method">getValue</span> gibt uns nur einen Integer zurück. Würden wir einen eigenen Datentyp zurückgeben wollen, müsste dieser <span class="interface">java.io.Serializable</span> implementieren.
	
</div>

### Implementierung

Als Nächstes wollen wir das zuvor angelegte Interface implementieren. Um uns ein bisschen Arbeit zu ersparen, erweitern wir einfach eine bereits von Java zur Verfügung gestellte Implementierung des Remote Interfaces, <span class="interface">java.rmi.server.UnicastRemoteObject</span>.

Unsere kleine Implementierung <span class="class">ServerImpl</span> könnte dann zum Beispiel so aussehen:

<pre class="code">import java.util.*;
import java.rmi.server.*;
import java.rmi.RemoteException;
public class ServerImpl extends UnicastRemoteObject implements Server
{
	private int value;
	
	public ServerImpl() throws RemoteException
	{
		this.value = 3;
	}

	public int getValue() throws RemoteException
	{
		return this.value;
	}
}</pre>

### Den Dienst zur Verfügung stellen

 Nun fehlt eigentlich nur noch, dass der Server dieses Server-Objekt auch wirklich zur Verfügung stellt. Dazu implementieren wir noch eine dritte Klasse , <span class="class">ServerProc</span> , die ansich nur eine einzige Aufgabe hat: Ein Server-Objekt in die RMIRegistry einzutragen und dieses Objekt somit auf einem festgelegten Port am Server zur Verfügung zu stellen.

Hierfür schreiben wir eine <span class="method">main</span>-Methode, die zuerst einen SecurityManager festlegt (das ist notwendig, da keine Remote-Objekte verwendet werden können, wenn keine SecurityManager festgelegt wurde) und danach ein <span class="class">Server</span>-Object in die RMIRegistry einträgt.

<pre class="code">import java.rmi.*;
import java.rmi.registry.*;
import java.rmi.server.UnicastRemoteObject;

class ServerProc
{
	public static void main (String args [])
	{
		if(System.getSecurityManager() == null)
		{
			System.setSecurityManager(new RMISecurityManager());
		}
		try {
			ServerImpl obj = new ServerImpl();
			Naming.rebind(&quot;rmi://localhost:1234/Server&quot;,obj);
		}catch (Exception e){
			[...]
		}
	
	}
}</pre>
				

Hiermit wird das Objekt "Server" auf Port 1234 auf localhost zur Verfügung gestellt.

### java.security

Wie bereits oben erwähnt, benötigt der Server einen SecurityManager um einen Dienst zur Verfügung stellen zu können. Über diesen SecurityManager kann festgelegt werden, welche Rechte "fremde" Klassen auf lokale Klassen haben. Um ein gewisses Mass an Sicherheit zu gewährleisten, muss jedes Programm, dass RMI verwendet, einen SecurityManager haben, ansonsten können nur Klassen aus dem lokalen classpath verwendet werden. [[ref-javasunrmitut][]]

Um hiermit nun Berechtigungen festzulegen, kann zum Beispiel eine java.policy-Datei angelegt werden.

Für unser Beispiel legen wir nur sehr einfache Berechtigungen fest, nämlich, dass es keine Einschränkungen der Verwendung gibt.

<div class="note">

Das hat nichts mit Variablen-/Klassen-/Methoden-Restriktionen wie "private", "public" oder "protected" zu tun, sondern geht noch viel tiefer. Es können sogar nur Zugriffe von bestimmten Netzwerkadressen eingerichtet werden.
	
</div>

<pre class="code">grant{
  permission java.security.AllPermission;
};</pre>

<div class="warning">
	
Die hiermit vergebenen Berechtigungen sind nur für dieses Beispiel gedacht. Für Details zur Vergabe <em>sinnvoller</em> Berechtigungen wird dringend empfohlen, die java.security-Dokumentation [<a href="http://java.sun.com/j2se/1.4.2/docs/api/java/security/package-summary.html">ref-javasec</a>] zu studieren.

</div>

Eine Einführung in java.security würde den Umfang dieses RMI-Tutorials bei weitem übersteigen, weshalb wir hiermit auf die java.security-Dokumentation [[ref-javasec](http://java.sun.com/j2se/1.4.2/docs/api/java/security/package-summary.html)] verweisen möchten.


### Kompilieren und Starten

Ansich kompiliert man RMI-Applikationen genau gleich wie jede andere Java-Applikation, jedoch mit einem kleinen Unterschied: Da die RMI-Applikation sehr weit weg von der eigentlich Netzwerkschicht geschrieben wird (keine Sockets usw.), gibt es mit rmic ein kleines Tool, das alle "fehlenden" Klassen erzeugt und kompiliert.

Nun wollen wir unseren kleinen Server einmal durchkompilieren:

<pre class="command">javac ServerProc.java
rmic ServerImpl</pre>
				

Wie man hier sieht, übergibt man rmic den Namen derjenigen Klasse, die den eigentlichen Serverdienst implementiert.

Nun müssen wir nur noch den Server starten. Hierzu benötigt RMI zunächst die rmiregistry, in der das zur Verfügung gestellte Objekt eingetragen wird. Da wir unseren Dienst auf Port 1234 anbieten wollen, starten wir rmiregistry so:

<pre class="command">rmiregistry 1234 &amp;</pre>
				

Jetzt können wir den Server starten, wobei wir die java.policy-Datei angeben müssen:

<pre class="command">java -Djava.security.policy=PolicyDatei ServerProc</pre>


## Die Client-Seite

Hier wollen wir nur einen Miniclient schreiben, der die <span class="method">getValue</span>-Methode des Server ausführt und ihren Rückgabewert ausgibt.

### Implementierung

<pre class="code">import java.rmi.*;
class Client
{			
	public static void main(String args [])
	{
		Server server;
		try {
			server = (Server)Naming.lookup(&quot;rmi://192.168.0.2:1234/Server&quot;);
			System.out.println(server.getValue());
		}catch(Exception e){
			[...]
		}
	}
}</pre>
				

Hier bindet der Client mit Hilfe der <span class="method">java.rmi.Naming.lookup</span>-Methode das Server-Objekt, das auf Port 1234 des Rechners mit der IP-Adresse 192.168.0.2 angeboten wird, in die lokale Variable server ein. Nun können die Methoden dieses Objektes ganz normal verwendet werden, was in der nächsten Zeile mit der Ausgabe auf den Bildschirm auch erfolgt.

<div class="note">
	
	Methoden des Remote-Objektes werden am Client immer nur aufgerufen jedoch nie direkt ausgeführt. Das wird spätestens dann ersichtlich, wenn wir die <span class="method">Server.getValue()</span>-Methode ein bisschen abändern, sodass sie etwas direkt auf den Monitor ausgibt:

<pre class="code">[...]
public int getValue() throws RemoteException
{
	System.out.println(&quot;getValue() invoked&quot;);
	return this.value;
}
[...]</pre>


	Jedesmal, wenn nun diese Methode am Client aufgerufen wird, wird am Server "getValue() invoked" ausgegeben.

</div>

### Kompilieren und Starten

 Zuerst kompilieren wir unseren Client mit...

<pre class="command">javac Client.java</pre>
				

Und führen ihn dann mit ...

<pre class="command">java Client</pre>
				

... aus. Es sollte nun eine "3" auf dem Bildschirm durch den Client ausgegeben werden.

Eine ausführliche Dokumentation von Java's RMI gibt es unter [http://java.sun.com/j2se/1.4.2/docs/api/](http://java.sun.com/j2se/1.4.2/docs/api/) .

[ref-javasunrmitut]: http://java.sun.com/docs/books/tutorial/rmi/TOC.html

## Autoren:

* [David Fritzer](mailto:dfritzer&ltat&gt;edu&lt;dot;gt;uni-klu.ac.at)
* [Horst Gutmann](mailto:hgutmann&ltat&gt;edu&lt;dot;gt;uni-klu.ac.at)
* [Sigrid Kuchler](mailto:skuchler&ltat&gt;edu&lt;dot;gt;uni-klu.ac.at)
