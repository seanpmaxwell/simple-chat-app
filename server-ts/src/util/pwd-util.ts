import bcrypt from 'bcrypt';


// Constants
const pwdSaltRounds = 12;


/**
 * Hash the password.
 */
function encrypt(password: string): Promise<string> {
    return bcrypt.hash(password, pwdSaltRounds);
}

/**
 * Hash the password synchronously. Useful for testing.
 */
function encryptSync(password: string): string {
    return bcrypt.hashSync(password, pwdSaltRounds);
}

/**
 * See if a password passed.
 */
function verify(password: string, pwdHash: string): Promise<boolean> {
    return bcrypt.compare(password, pwdHash);
}


// Export default
export default {
    encrypt,
    encryptSync,
    verify,
};
