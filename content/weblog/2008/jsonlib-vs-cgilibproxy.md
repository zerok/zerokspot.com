---
date: '2008-02-13T12:00:00-00:00'
language: en
tags:
- java
title: Json-lib vs. CGLib Proxies - 1:0
---


Sometimes I'm really blind. For the last couple of months I was all the time having some problems with serializing JavaBeans that I get out of [Hibernate](http://www.hibernate.org) into [JSON](http://json.org) and often got an error like this:

    java.lang.NoSuchMethodException: Property 'delegate' has no getter method
    

I always knew that it had to be something with the whole proxy-layer Hibernate uses for handling for instance lazy referencing, but for some reason I never really actually did anything against it but rewrite the code to build the JSONObjects more or less manually.

-------------------------------

Explicitly setting lazy="false" in the respective ForeignKey-column in the mapping file resolves this, but what if I want some laziness in my model and simply not serialize what's behind that reference?

After posting on the [Json-lib](http://json-lib.sourceforge.net/) mailing-list and following Andres Almiray's hint of taking a look at [PropertyFilter](http://json-lib.sourceforge.net/apidocs/jdk15/net/sf/json/util/PropertyFilter.html), I wrote another one (I wrote one months ago when I first got this problem but filtered for the name "delegate"), still with no luck. Then I took a look at the source code of the Json-lib and noticed, that the PropertyFilter gets applied *after* all the properties are extracted using [beanutils](http://commons.apache.org/beanutils/ "Commons - BeanUtils")' PropertyUtil, which is actually where the exception is thrown in the first place. So back to the source code and a few lines up ;-) Given that I'm already inside the bean, I thought the BeanProcessor must be where I actually want to add my modifications. And right I was:

So basically all you need is a custom [JsonBeanProcessor](http://json-lib.sourceforge.net/apidocs/jdk15/net/sf/json/processors/JsonBeanProcessor.html) that returns nothing but an empty JSONObject, 

@@ java @@
import net.sf.json.JSONObject;
import net.sf.json.JsonConfig;
import net.sf.json.processor.JsonBeanProcessor;

public class HibernateJsonBeanProcessor implements JsonBeanProcessor {

	public JSONObject processBean(Object obj, JsonConfig jsonConfig) {
		return new JSONObject();
	}

}
@@

a custom [JsonBeanProcessorMatcher](http://json-lib.sourceforge.net/apidocs/jdk15/net/sf/json/processors/JsonBeanProcessorMatcher.html) that maps the custom processor from above to the offending Bean, 

@@ java @@
import net.sf.json.processor.JsonBeanProcessorMatcher;
import java.util.Set;
import org.apache.log4j.Logger;

public class HibernateJsonBeanProcessorMatcher extends JsonBeanProcessorMatcher {
	
	private static Logger log = Logger.getLogger(HibernateJsonBeanProcessorMatcher.class);
	
	@Override
	public Object getMatch(Class target, Set set) {
		if (target.getName().contains("$$EnhancerByCGLIB$$")) {
			log.warn("Found Lazy-References in Hibernate object "
					+ target.getName());
			return org.hibernate.proxy.HibernateProxy.class;
		}
		return DEFAULT.getMatch(target, set);
	}

}
@@

and last but not least the [JsonConfig](http://json-lib.sourceforge.net/apidocs/jdk15/net/sf/json/JsonConfig.html) that ties it all together:

@@ java @@
JsonConfig conf = new JsonConfig();
conf.registerJsonBeanProcessor(org.hibernate.proxy.HibernateProxy.class, 
    new HibernateJsonBeanProcessor());
conf.setJsonBeanProcessorMatcher(new HibernateJsonBeanProcessorMatcher());
@@
