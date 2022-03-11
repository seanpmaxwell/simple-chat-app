// Wait for a certain number of milliseconds.
export function tick(milliseconds: number): Promise<void> {
    return new Promise((resolve) => {
        setTimeout(() => {
            resolve();
        }, milliseconds);
    });
}
