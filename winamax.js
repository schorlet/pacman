// card values
var values = {
    "2": 2,
    "3": 3,
    "4": 4,
    "5": 5,
    "6": 6,
    "7": 7,
    "8": 8,
    "9": 9,
    "10": 10,
    "J": 11,
    "Q": 12,
    "K": 13,
    "A": 14
};

// extract value from card
function value(card) {
    var val = card.slice(0, -1);
    return values[val];
}

// read player1 cards
var n = parseInt(readline(), 10),
    player1 = [];
for (var x = 0; x < n; x++) {
    player1.push(value(readline()));
}

// read player2 cards
var m = parseInt(readline(), 10),
    player2 = [];
for (var x = 0; x < m; x++) {
    player2.push(value(readline()));
}

var ply1 = [],
    ply2 = [];

var battle = false,
    battle_count = 0;


while (player1.length > 0 && player2.length > 0) {
    ply1.push(player1.shift());
    ply2.push(player2.shift());

    if (!battle) {
        battle_count += 1;
    }

    if (ply1[ply1.length - 1] == ply2[ply2.length - 1]) {
        battle = true;
        if (player1.length < 3 || player2.length < 3) {
            break;
        }
        ply1.push.apply(ply1, player1.splice(0, 3));
        ply2.push.apply(ply2, player2.splice(0, 3));

    } else {
        battle = false;
        if (ply1[ply1.length - 1] > ply2[ply2.length - 1]) {
            player1.push.apply(player1, ply1.concat(ply2));
        } else {
            player2.push.apply(player2, ply1.concat(ply2));
        }
        ply1.splice(0);
        ply2.splice(0);
    }
}

if (battle) {
    print("PAT");
} else {
    var winner_id = (player1.length > 0) ? 1 : 2;
    print(winner_id + " " + battle_count);
}


