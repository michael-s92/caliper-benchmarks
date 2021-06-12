'use strict';

const docType = "election-doc";

class Election {

    constructor(id, candidates, votes, isClosed){
        this.docType = docType;
        this.id = id;
        this.candidates = (candidates === undefined) ? [] : candidates;
        this.votes = (votes === undefined) ? [] : votes;
        this.isClosed = (isClosed === undefined) ? false : votes;
    }

    static fromJSON(obj){
        if(obj.id !== undefined && obj.candidates !== undefined && obj.votes !== undefined){
            return new Election(obj.id, obj.candidates, obj.votes, obj.isClosed);
        }
    }
}

module.exports = Election