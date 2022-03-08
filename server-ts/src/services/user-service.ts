import userRepo from '@repos/user-repo';
import User, { IUser, IUserCreds } from '@models/user-model';
import pwdUtil from '@util/pwd-util';


// **** Types/Constants **** //

// Errors
const errors = {
    addOne: 'User cound not be saved',
} as const;


// **** Functions **** //

/**
 * Add one user.
 */
async function addOne(email: string, name: string, password: string): Promise<void> {
    // Save the user
    const newUser = User.new(email, name);
    const newUserId = await userRepo.addOne(newUser);
    if (!newUserId) {
        throw Error(errors.addOne);
    }
    newUser.id = newUserId;
    // Once we have we have the new user's id, insert the password
    const pwdHash = await pwdUtil.encrypt(password);
    const creds: IUserCreds = {
        pwdHash,
        userId: newUser.id,
    };
    await userRepo.addCreds(creds);
}

/**
 * Fetch all users.
 */
function fetchAll(): Promise<IUser[]> {
    return userRepo.fetchAll();
}


// Export default
export default {
    addOne,
    fetchAll,
} as const;
