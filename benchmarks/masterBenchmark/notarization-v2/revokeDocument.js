/* eslint-disable no-undef */
'use strict';

const utils = require('./utils');
const seeds = require('./seeds.json');

class revokeDocument {

    static get() {

        // select random document from initDocuments
        let randomAccessKey = 0;
        do {
            randomAccessKey = utils.getRandomInt(seeds.initDocuments.length);
        } while (seeds.initDocuments[randomAccessKey] === undefined);

        let doc = seeds.initDocuments[randomAccessKey];

        // revokeDocument(ctx, documentKey, custodianKey) 

        let args = {
            chaincodeFunction: 'revokeDocument',
            chaincodeArguments: [doc.documentId, doc.custodian.key]
        };


        return args;

    }
}

module.exports = revokeDocument;
