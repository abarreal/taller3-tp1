function randint(min, max) {
    min = Math.ceil(min);
    max = Math.floor(max);
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

async function simulateActivity(mean, variation) {
    // Generate a random amount of milliseconds.
    const delay = randint(mean-variation, mean+variation)/2
    const start = new Date()

    while (true) {
        // Compute the time difference in milliseconds since the start.
        const waited = (new Date().getTime() - start.getTime());
        if (waited > delay) {
            break;
        }
    }
    // Block.
    await new Promise(r => setTimeout(r, delay));
}

module.exports = {
    randint: randint,
    simulateActivity: simulateActivity,
}