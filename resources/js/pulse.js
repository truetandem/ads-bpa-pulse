$(document).ready(function(){
    var data = [0,     0,     0,    0,    0,    0,     0,     0,     0,    0,
                0,     0,     0,    0,    0,    0,     0,     0,     0,    0,
                0,     0,     0,    0,    0,    0,     0,     0,     0.08, 0.18,
                0.08,  0,     0,    0,    0,    0,     0,    -0.04, -0.08, 0.3,
                0.7,   0.3,  -0.17, 0.00, 0.04, 0.04,  0.05,  0.05,  0.06, 0.07,
                0.08,  0.10,  0.11, 0.11, 0.10, 0.085, 0.06,  0.04,  0.03, 0.01,
                0.01,  0.01,  0.01, 0.02, 0.03, 0.05,  0.05,  0.05,  0.03, 0.02,
                0,     0,     0];

    var idx = 0;
    var lastData = 0;

    var ecg = new PlethGraph("ecg", function () {
        if (idx++ >= data.length - 1) {
            idx=0;
        }

        // Create a random function that is dependent on the last value
        var hysteresisRandom = function() {
            lastData += (Math.floor((Math.random() * 5) + 1) - 3) / 50;
            if (Math.abs(lastData) >= 1) lastData = (lastData > 0) ? 1 : -1;
            return lastData;
        }

        return [data[idx] + hysteresisRandom() / 10];
    });

    ecg.speed = 1.5;
    ecg.scaleFactor = 0.8;
    ecg.start();
});
