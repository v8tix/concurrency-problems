bool mu1 = false;
bool mu2 = false;

proctype proc1() {
    atomic {
        mu1 = true;
        mu2 = true;
    }
    mu1 = false;
    mu2 = false;
}

proctype proc2() {
    atomic {
        mu2 = true;
        mu1 = true;
    }
    mu2 = false;
    mu1 = false;
}

init {
    run proc1();
    run proc2();
}
