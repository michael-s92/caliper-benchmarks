/* eslint-disable no-undef */
'use strict';

const utils = require('./utils');
const seeds = require('./seeds.json');
const Utils = require('../../../src/masterCC/eureka/lib/utils');

class submittingArticle {

    static get() {
	    let args;

        let randomAccessKey = 0;
        let articles = seeds.newArticles;

        do{
            randomAccessKey = utils.getRandomInt(articles.length);
        } while(articles[randomAccessKey] === undefined);

        let article = articles[randomAccessKey];

        let randomTitle = "A" + Utils.getRandomInt(100) + ": " + article.title;

        // submittingArticle(ctx, title, author_id, coauthor_ids, ref_author_ids, fee, lref, authorKey)
	    args = {
                chaincodeFunction: 'submittingArticle',
                chaincodeArguments: [randomTitle, article.author.id, JSON.stringify(article.coauthor_ids), JSON.stringify(article.refauthor_ids), article.fee, article.lref, article.author.key]
            };

	    return args;

	}
}

module.exports = submittingArticle;
