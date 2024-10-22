"use strict";
/**
 * This would have been much simpler in AE, but I wanted to mix it up and see what animating paths was like
 */
const ballName = "ball";
const cloudNames = [
    "cloud0",
    "cloud1",
    "cloud2",
    "cloud3",
    "cloud4",
    "cloud5",
    "cloud6",
    "cloud7"
];
const smokeNames = ["smoke0", "smoke1", "smoke2"];
const cannonNames = ["cannon0", "cannon1", "cannon2", "cannon3", "cannon4"];
const createBallTarget = (cannonIndex) => `[data-name="${cannonNames[cannonIndex]}${ballName}"]`;
const createSmokeTarget = (cannonIndex, smokeIndex) => `[data-name="${cannonNames[cannonIndex]}${smokeNames[smokeIndex]}"]`;
const createCloudTarget = (cloudIndex) => `[data-name="${cloudNames[cloudIndex]}"]`;
const randomIntFromInterval = (min, max) => Math.floor(Math.random() * (max - min + 1) + min);
const randomDoubleFromInterval = (min, max) => (Math.random() * (max - min + 1) + min).toFixed(2);
// used to make transformOrigin be center in reference to the path rather than the entire SVG
const setTransformOrigin = (target) => {
    const path = document.querySelectorAll(target)[0];
    const { x, y, height, width } = path.getBBox();
    path.style.transformOrigin = `${x + width / 2}px ${y + height / 2}px`;
};
const resetCannon = (cannonIndex) => {
    anime.set(createBallTarget(cannonIndex), {
        opacity: 0,
        translateX: 0,
        translateY: 0
    });
    setTransformOrigin(createCloudTarget(2));
    smokeNames.forEach((st, i) => {
        setTransformOrigin(createSmokeTarget(cannonIndex, i));
        anime.set(createSmokeTarget(cannonIndex, i), {
            opacity: 0,
            translateX: 0,
            translateY: 0
        });
    });
};
const fireCannon = (cannonIndex) => {
    const smokeIndex = randomIntFromInterval(0, 2);
    const smokeTarget = createSmokeTarget(cannonIndex, smokeIndex);
    const ballTarget = createBallTarget(cannonIndex);
    const { height, width, x, y } = document
        .querySelectorAll(smokeTarget)[0]
        .getBBox();
    let smokeTimeLine = anime.timeline({
        easing: "easeOutElastic(1, .8)"
    });
    smokeTimeLine.set({
        targets: smokeTarget,
        rotate: `${randomIntFromInterval(0, 180)}deg`,
        transformOrigin: `${x}px ${y}px 0`
    });
    smokeTimeLine
        .add({
        targets: smokeTarget,
        opacity: [0.0, 1],
        scale: [1, randomDoubleFromInterval(2.5, 4)],
        duration: 1000
    })
        .add({
        targets: smokeTarget,
        opacity: [1, 0],
        duration: 2000
    }, "-=700");
    let cannonBallTimeLine = anime.timeline({ easing: "linear" });
    cannonBallTimeLine
        .add({
        targets: ballTarget,
        opacity: [0.0, 1],
        duration: 20
    })
        .add({
        targets: ballTarget,
        translateX: `-${randomIntFromInterval(31, 33)}%`,
        translateY: [
            {
                value: -1 * randomIntFromInterval(20, 30),
                easing: "easeOutCubic"
            },
            {
                value: randomIntFromInterval(62, 64),
                easing: "easeInCirc"
            }
        ],
        complete: () => anime({
            targets: createCloudTarget(2),
            scale: [1, 1.05, 1],
            duration: 250,
            easing: `spring(1, 80, 10, 20)`
        })
    })
        .add({
        targets: ballTarget,
        opacity: [1, 0.0],
        duration: 20
    });
};
const resetCannons = () => {
    resetCannon(0);
    resetCannon(1);
    resetCannon(2);
    resetCannon(3);
    resetCannon(4);
};
resetCannons();
const cloudAnimation = (i, maxDelta) => {
    anime({
        targets: createCloudTarget(i),
        translateX: () => anime.random(-1 * maxDelta, maxDelta),
        translateY: () => anime.random(-1 * maxDelta, maxDelta),
        duration: randomIntFromInterval(3000, 5000),
        easing: `easeInOutSine`,
        complete: () => cloudAnimation(i, maxDelta)
    });
};
const animateSurroundingClouds = () => {
    cloudAnimation(0, 8);
    cloudAnimation(1, 8);
    cloudAnimation(2, 8);
    cloudAnimation(3, 8);
};
const animateShipClouds = () => {
    cloudAnimation(4, 8);
    cloudAnimation(5, 8);
    cloudAnimation(6, 8);
    cloudAnimation(7, 8);
};
animateSurroundingClouds();
animateShipClouds();
const staggerFire = (i) => {
    setTimeout(() => fireCannon(i), i * 200);
};
const setupReoccuringFire = (i) => {
    setTimeout(() => {
        resetCannon(i);
        fireCannon(i);
        setupReoccuringFire(i);
    }, randomIntFromInterval(3000, 5000));
};
setTimeout(() => {
    resetCannons();
    staggerFire(0);
    staggerFire(1);
    staggerFire(2);
    staggerFire(3);
    staggerFire(4);
    setupReoccuringFire(0);
    setupReoccuringFire(1);
    setupReoccuringFire(2);
    setupReoccuringFire(3);
    setupReoccuringFire(4);
}, 3000);