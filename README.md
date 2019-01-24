# Schemas demo

This demo illustrates the proposal found [in this doc](https://docs.google.com/document/d/1sLWGZR4kCvQVGv-TA5x8ny-AxL3ChBYNeFYW1eACsDw/).  It contains a hypothetical "global" schema, and an individual "nihms" schema. 

## Prerequisites

* `npm` needs to be installed
* You need to globally install an npm static http server via `npm install http-server -g`
* You need to install browserify via `npm install -g browserify`

## Instructions

The following needs to only be done once:

* do `npm install` to gather all the js reqirements
* do `browserify init.js -o scripts/bundle.js`.  This will create scripts/bundle.js containing a JSON schema validator, and a JSON reference dereferencer.

To run the demo:

* do `http-server`

## The demo

Point your browser to `http://localhost:8080/forms/nihms.html`

This renders the an Alpaca form (located in `forms/hihms.html`) based on the schema `schemas/nihms.json`.  Look at the html file, ant try it!

When you enter something in the NLMTA field and click the "Send Form Data" button,a popup will appear, showing validation status.  The demo starts with a metadata blob that is *invalid* with respect to the global schema:  it contains a field "foo" not defined in the global schema, and it is also missing required fields.

Play with the metadata blob (present in forms/nihms.html inline javascript), or the schemas.  Edit the files and reload the brrowser.

Can you:

* Edit the metadata blob so that it'll pass validation after submitting the form?
* Edit the nihms schema to remove global schema validation, or the additional validation prerequisites defined in that schema?
