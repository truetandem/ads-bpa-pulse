var lastData = 0;

// Create a random function that is dependent on the last value
function hysteresisRandom(){
    lastData += (Math.floor((Math.random() * 5) + 1)-3)/50;
    if (Math.abs(lastData) >= 1) lastData = (lastData > 0) ? 1 : -1;
    return lastData;
}

// Generate a real time data grab of various length
function generateData(){
    buffer = new Array();
    var inputLength = Math.floor((Math.random() * 1) + 1);;
    for( i = 0 ; i < inputLength ; i++ ) buffer[i] = hysteresisRandom();
    return buffer;
}

var ECG_data = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 
                0.08, 0.18, 0.08, 0, 0, 0, 0, 0, 0, -0.04, 
                -0.08, 0.3, 0.7, 0.3, -0.17, 0.00, 0.04, 0.04, 
                0.05, 0.05, 0.06, 0.07, 0.08, 0.10, 0.11, 0.11, 
                0.10, 0.085, 0.06, 0.04, 0.03, 0.01, 0.01, 0.01, 
                0.01, 0.02, 0.03, 0.05, 0.05, 0.05, 0.03, 0.02, 0, 0, 0];

var ECG_idx = 0;

function getECG(){
    if (ECG_idx++ >= ECG_data.length - 1) ECG_idx=0;
    var output = new Array();
    output[0] = ECG_data[ECG_idx] + hysteresisRandom()/10;
    return output;
}
var ecg;

$(document).ready(function(){
    ecg = new PlethGraph("ecg", getECG);
    ecg.speed = 1.5;
    ecg.scaleFactor = 0.8;
    ecg.start();
});
