import bcrypt from 'bcrypt';


// **** Vals **** //

// Misc
const pwdSaltRounds = 12;


// **** Functions **** //

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
