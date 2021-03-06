---
date: '2013-09-15T16:30:01+02:00'
language: en
tags:
- weeklyresearch
- json
- jsonschema
title: 'Weekly Research: JSON-Schema'
---


Working with XML in the Java world is made easy by one thing: most applications provide you with a set of clear rules for how that XML should look like and there are tools to convert these standardized definitions into executable code.

If you work with JSON on the other hand I've become used to just looking at code examples and hoping that their author didn’t forget any detail. So, what’s lacking in my tool-chain right now is something that helps me validate JSON against something like a schema (comparable to an XSD or DTD for XML). I also sometimes have to design RESTful APIs that serve JSON content. In order to better communicate what the served data should look like and also to provide some test-cases myself, something like that would be immensely useful.

After looking for quite a while, the only definition system I could find was [JSON-Schema][schema] by [Kris Zyp](http://www.sitepen.com/blog/author/kzyp/) and [Gary Court](https://github.com/garycourt). I was a little surprised since I had expected to find some competition in that field similar to what had happened back when XML was still new and every large player had their own little schema definition and validation framework. But since I don’t have all that much time this week, this is just perfect :-)

-----------

## JSON-Schema

JSON is extremely simple compared to XML and this also shows in the length of the JSON-Schema specification. You can easily get through it within an hour or so, and this is definitely a good thing :-)

The whole specification is split into 3 documents:

* [Core][s1]
* [Validation][s2]
* [Hyper-Schema][s3]

The first document describes things like the interaction of the spec with things like HTTP, MIME, and JSON in general while the second includes the actual schema structure.

The third specification of the JSON-Schema-bundle is all about defining JSON for hypermedia including things like linking and media-types. This is a bit of a weird one since there is definitely some overlapping here with Web-Service/API specifications like [WADL][] and [JSON-API][jsonapi] here. Because of this I will skip this document here.

### Basic structure

To keep this as pratical as possible, let’s work on a small example document. Let’s say you want to represent a simple shopping cart implementation in JSON. It needs a total price, items with product, price and quantity information, vouchers and shipping costs:

```json
{
    "total": 170.0,
    "currency": "EUR",
    "items": [
        {
            "quantity": 1,
            "basePrice": 100.0,
            "code": "prod1",
            "name": "Table Martha",
            "price": 100
        },
        {
            "quantity": 4,
            "basePrice": 20.0,
            "code": "chair1",
            "name": "Chair Eddy",
            "price": 80
        }
    ],
    "vouchers": [
        {
            "name": "Dinner discount",
            "value": 20,
            "type": "absolute",
            "code": "dindis"
        }
    ],
    "shippingCosts": 10.0
}
```

So let’s definite a schema for this by starting with the defining that we are working with an object with the properties mentioned above:

<pre><code class="language-json">{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "title": "Cart",
    "descrption": "A simple cart abstraction",
    "type": "object",
    "properties": {
        "total": {},
        "currency": {},
        "items": {},
        "shippingCosts": {},
        "vouchers": {}
    }
}</code></pre>

The “$schema” property is mandatory and always has to be at the root-level of the schema definition. It indicates that what you have here really is a JSON-Schema and also what version of it is used. The specification defines a couple of keywords (like “type” and “properties”) and if you use custom keywords the $schema has to point to your extension’s specification.

Before going into every property to define what format it should have, let’s make sure that all the required properties are always there and that no additional properties are allowed by adding this to the root of the schema:

<pre><code class="language-json">"required": ["total", "currency", "items", "shippingCosts", 
             "vouchers"],
"additionalItems": false</code></pre>

But now, on to the “simple” values here:

<pre><code class="language-json">"total": {
    "type": "number",
    "minimum": 0
},
"currency": {
    "enum": ["EUR"]
},
"shippingCosts": {
    "type": "number",
    "minimum": 0
}</code></pre>

As we’ve already seen in the root-object, there is a “type” keyword that specifies the value’s type we can expect. For “total” and “shippingCosts” we want to restrict the possible value to be numeric (hence “number”) and positive or zero (hence setting the [minimum][nummin] value to “0”). JSON-Schema defines a handful [types][] that you should already know by dealing with JSON itself: “array”, “boolean”, “integer”, “number”, “null”, “object”, and “string”.

For the currency property we already know what specific values it accepts so we can put them all into an “enum”.

“vouchers” and “items” are a bit more complex since they are arrays of objects. Because the voucher object is a bit shorter, let’s stick with it for now:

<pre><code class="language-json">"vouchers": {
    "type": "array",
    "items": {
        "type": "object",
        "required": ["name", "value", "type", "code"],
        "properties": {
            "name": {
                "type": "string",
                "minimumLength": 1
            },
            "value": {
                "type": "number",
                "minimum": 0
            },
            "type": {
                "enum": ["relative", "absolute"]
            },
            "code": {
                "type": "string",
                "pattern": "^[a-z0-9]+$"
            }
        }
    }
}</code></pre>

Here we define that “vouchers” is an array containing items that have a name, value, type and code. New is the property “pattern” with which you can restrict string values to a certain regular expression.

### References & definition-reuse

I’ll skip the “items” definition here because of its similarity to vouchers. In fact, there are quite a few common definitions here which would be awesome to be reused. The name of products and vouchers have the same structure and prices also look the same. Thankfully, there is a way to create a single definition and then reference it from where you need it. So, this is the final schema for our cart that also includes references:

<pre><code class="language-json">{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "title": "Cart",
    "descrption": "A simple cart abstraction",
    "type": "object",
    "properties": {
        "total": {"$ref": "#/definitions/price"},
        "currency": {
            "enum": ["EUR"]
        },
        "items": {
            "type": "array",
            "items": {"$ref": "#/definitions/item"}
        },
        "shippingCosts": {"$ref": "#/definitions/price"},
        "vouchers": {
            "type": "array",
            "items": {"$ref": "#/definitions/voucher"}
        }
    },
    "required": [
        "total", "currency", "shippingCosts", "items"
    ],
    "additionalProperties": false,
    "definitions": {
        "price": {
            "type": "number",
            "minimum": 0
        },
        "name": {
            "type": "string",
            "minimumLength": 1
        },
        "code": {
            "type": "string",
            "pattern": "^[a-z0-9]+$"
        },
        "item": {
            "type": "object",
            "properties": {
                "name": {"type": "string", "minimumLength": 1},
                "code": {"$ref": "#/definitions/code"},
                "quantity": {"type": "integer", "minimum": 1}
            }
        },
        "voucher": {
            "type": "object",
            "required": ["name", "value", "type", "code"],
            "properties": {
                "name": { "$ref": "#/definitions/name" },
                "value": {
                    "type": "number",
                    "minimum": 0
                },
                "type": {"enum": ["relative", "absolute"]},
                "code": { "$ref": "#definitions/code" }
            }
        }
    }
}</code></pre>

The examples in the specifications always group reusable definitions into a “definitions” section so I opted to do the same here. Referencing itself is based around URIs with the document acting as initial scope. Some parts of this are optional for implementations to support. You can find more about this in the [“URI resolution scopes and dereferencing”](http://json-schema.org/latest/json-schema-core.html#anchor25) chapter of the specification.

### Interdependent definitions

If you’re working with more complex structures that change depending on the value of a single property, you will probably have to resort to use something like the “oneOf” property which says that the data should match one of many definitions. You can find an [example](http://json-schema.org/example2.html) on json-schema.org that creates an fstab-representation in JSON.

### Referencing a Schema

In XML you can reference the schema/DTD a document should adhere to either within the DOCTYPE preamble or as part of the namespace definition within the opening `<?xml ?>` statement. JSON doesn’t offer anything like that so the [JSON-Schema specification](http://json-schema.org/latest/json-schema-core.html#anchor33) opted to put this metadata into the transport layer and HTTP in specific.

Two methods have been defined:

* Extend the Content-Type of the document with a profile-parameter to reference the schema’s URI: `Content-Type: application/json;profile=http://example.org/schema.json#`
* ... or use the [Link-header][linkhdr] in combination with `rel=describedBy`: `Link: <http://example.org/schema.json#>; rel="describedBy"`

In either case the used URI *must* point to an actual JSON-Schema document.

## Validators & Tools

A schema language doesn’t really help if you don’t have tools to work with it. Luckily, the spec’s website already [lists quite a few projects](http://json-schema.org/implementations.html) offering solutions for validating, generating and much more. For this article I’ve played around with Julian Berman’s [jsonschema][py] library for Python in order to validate my examples. Here is the small validator script I used:

```python
#!/usr/bin/env python
import argparse
import jsonschema
import json

argp = argparse.ArgumentParser()
argp.add_argument('schema')
argp.add_argument('file')
opts = argp.parse_args()

schema = None
data = None

with open(opts.schema) as fp:
    schema = json.load(fp)
with open(opts.file) as fp:
    data = json.load(fp)

jsonschema.validate(data, schema)
```

## Conclusion

JSON-Schema definitely looks like something I’ll be using in the future. I have, for instance, a set of API tests for making sure that an API doesn’t change while I’m working on tools using it. Right now the test-cases are very simple and this definition language might help extend the scope of these tests while not increasing their complexity.

[s1]: http://json-schema.org/latest/json-schema-core.html
[s2]: http://json-schema.org/latest/json-schema-validation.html
[s3]: http://json-schema.org/latest/json-schema-hypermedia.html
[schema]: http://json-schema.org/
[dw]: http://davidwalsh.name/json-validation
[jsonapi]: http://jsonapi.org/
[wadl]: http://www.w3.org/Submission/wadl/
[linkhdr]: http://tools.ietf.org/html/rfc5988#section-6.1
[nummin]: http://json-schema.org/latest/json-schema-validation.html#anchor21
[types]: http://json-schema.org/latest/json-schema-core.html#anchor8
[py]: https://github.com/Julian/jsonschema
