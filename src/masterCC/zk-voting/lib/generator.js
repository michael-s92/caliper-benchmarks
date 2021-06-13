'use strict';

const fs = require('fs');
const read = require('read-yaml');
const RandExp = require('randexp');

const Utils = require('./utils');
const { title } = require('process');


const parameters = read.sync('seedParameters.yaml');

let admins = [];
let initElections = [];
let newElections = [];

function generateAdmin(index) {

    let id = Utils.generateRandomWord(parameters.id_length) + index;
    let key = Utils.generateRandomKey(parameters.key_length);
    let name = "A: " + parameters.names[Utils.getRandomInt(parameters.names.length)];

    return {
        id: id,
        name: name,
        key: key
    };
}

function generateElection(i, flag){

    let id = flag + i + ": " + Utils.generateRandomKey(parameters.id_length);
    let len = parameters.min_candidate_per_election + Utils.getRandomInt(parameters.max_candidate_per_election - parameters.min_candidate_per_election);
    let candidates = Utils.getRandomArrayOfStrings(len);

    return {
        id: id,
        candidates: candidates
    };
}


for (let i = 0; i < parameters.admins; i++) {
    admins.push(generateAdmin(i));
}

for (let i = 0; i < parameters.init_elections; i++) {
    initElections.push(generateElection(i, 'Init'));
}

for (let i = 0; i < parameters.new_elections; i++) {
    newElections.push(generateElection(i, 'New'));
}

const json = JSON.stringify({
    admins: admins,
    initElections: initElections,
    newElections: newElections
}, null, 4);


fs.writeFile('seeds.json', json, function (err) {
    if (err) {
        console.log(err);
    }
});

//console.log("=============================== Generate seeds.json done");
