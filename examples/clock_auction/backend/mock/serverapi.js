//  SJ: install express
var express = require('express');
var app = express();

app.use(express.json());

app.use(function (req, res, next) {
    res.header("Access-Control-Allow-Origin", "*");
    res.header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept");
    res.header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS");
    next();
});

app.get('/',(function(req,res){
    res.send('Welcome to Spectrum Auction by Elbonia Communication Commission - Enabled by FPC');
}));

//  route for getAuctionDetails
app.get('/getAuctionDetails',(function(req,res)
{
    var returnStr;
    var jsonAuction = {
        "auctionName": "Spectrum Auction 2020",
        "owner": "Elbonian Communications Commission",
        "clockPriceIncrease": "5000",
        "currentPhaseStatus": "1",
        "currentRoundNumber": "1"
    };

    //  Add other fields
    jsonAuction["territories"] = [
        {
            "name": "Elbograd",
            //"comment": "encumbrance for each license; License id = index of the license in array",
            "numLicenses": 4,
            "licenses": [
                {"licenseId":1,"encumbrance": 0},
                {"licenseId":2,"encumbrance": 10},
                {"licenseId":3,"encumbrance": 25},
                {"licenseId":4,"encumbrance": 45},
            ],
            "minimumPrice": "10000"
        },
        {
            "name": "Mudberg",
            "numLicenses": 5,
            "licenses": [
                {"licenseId":1,"encumbrance": 0},
                {"licenseId":2,"encumbrance": 10},
                {"licenseId":3,"encumbrance": 25},
                {"licenseId":4,"encumbrance": 25},
                {"licenseId":5,"encumbrance": 45},
            ],
            "minimumPrice": "6000"
        },
        {
            "name": "Deserton",
            "numLicenses": 3,
            "licenses": [
                {"licenseId":1,"encumbrance": 0},
                {"licenseId":2,"encumbrance": 10},
                {"licenseId":3,"encumbrance": 25},
            ],
            "minimumPrice": "7000"
        },
        {
            "name": "Phlimsk",
            "numLicenses": 6,
            "licenses": [
                {"licenseId":1,"encumbrance": 0},
                {"licenseId":2,"encumbrance": 10},
                {"licenseId":3,"encumbrance": 25},
                {"licenseId":4,"encumbrance": 45},
                {"licenseId":5,"encumbrance": 45},
                {"licenseId":6,"encumbrance": 45},
            ],
            "minimumPrice": "8000"
        }];

    jsonAuction["bidders"] = [{
        "name": "A-Telco"
    }, {
        "name": "B-Telco"
    }, {
        "name": "C-Telco"
    }];
    returnStr = JSON.stringify( jsonAuction);

    res.send(returnStr);
}));

//  route for getRoundResults
app.get('/getRoundResults',(function(req,res)
{
    res.send("getRoundResults");
}));


console.log ("Listening on localhost:3000...");
var server=app.listen(3000,function() {});

