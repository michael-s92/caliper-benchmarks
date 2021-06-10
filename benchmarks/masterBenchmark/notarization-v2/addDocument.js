/* eslint-disable no-undef */
'use strict';

const utils = require('./utils');
const seeds = require('./seeds.json');

class addDocument {

    static get() {

        let randomAccessKey = 0;
        let docs = seeds.benchmarkDocuments;

        do{
            randomAccessKey = utils.getRandomInt(docs.length);
        } while(docs[randomAccessKey] === undefined);

        let doc = docs[randomAccessKey];


        // addDocument(ctx, documentHash, custodianId, custodianKey, studentId, studentKey)

	    let args = {
                chaincodeFunction: 'addDocument',
                chaincodeArguments: [doc.documentHash, doc.custodian.id, doc.custodian.key, doc.student.id, doc.student.key]
            };

	    return args;

	}
}

module.exports = addDocument;
