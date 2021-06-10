/* eslint-disable no-undef */
'use strict';

const utils = require('./utils');
const seeds = require('./seeds.json');

class getAllDocumentReaders {

    static get() {

        // select random document from initDocuments
        let randomAccessKey = 0;
        do {
            randomAccessKey = utils.getRandomInt(seeds.initDocuments.length);
        } while (seeds.initDocuments[randomAccessKey] === undefined);

        let doc = seeds.initDocuments[randomAccessKey];

        // getAllDocumentReaders(ctx, documentKey, studentKey)

        let args = {
            chaincodeFunction: 'getAllDocumentReaders',
            chaincodeArguments: [doc.documentId, doc.student.key]
        };


        return args;

    }
}

module.exports = getAllDocumentReaders;
