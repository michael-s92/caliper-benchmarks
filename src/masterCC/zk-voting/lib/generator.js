'use strict';

const fs = require('fs');
const read = require('read-yaml');
const RandExp = require('randexp');

const Utils = require('./utils');
const { title } = require('process');


const parameters = read.sync('seedParameters.yaml');

let admins = [];

function generateAdmin(type, index) {

    let id = type + Utils.generateRandomWord(parameters.id_length) + index;
    let key = Utils.generateRandomKey(parameters.key_length);
    let name = type + ": " + parameters.names[Utils.getRandomInt(parameters.names.length)];

    return {
        id: id,
        name: name,
        key: key
    };
}




for (let i = 0; i < parameters.admins; i++) {
    admins.push(generateAdmin(authorUserType, i));
}

const json = JSON.stringify({
    admins: admins

}, null, 4);

fs.writeFile('seeds.json', json, function (err) {
    if (err) {
        console.log(err);
    }
});

//console.log("=============================== Generate seeds.json done");
