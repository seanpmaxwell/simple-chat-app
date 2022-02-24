import knex from './knex';
import { IUser, IUserCreds, TSavedUser } from '@models/user-model';


/**
 * Find by email.
 * 
 * @param email 
 * @returns 
 */
async function findByEmail(email: string): Promise<TSavedUser | null> {
    const resp = await knex<IUser, TSavedUser[]>('users').where('email', email);
    return (resp.length > 0 ? resp[0] : null);
}


/**
 * Fetch the password hash.
 * 
 * @param userId 
 */
async function fetchPwdHash(userId: number): Promise<string> {
    const resp = await knex<IUserCreds, IUserCreds[]>('userCreds').where('userId', userId);
    return (resp.length > 0 ? resp[0].pwdHash : '');
}


/**
 * Add one user.
 * 
 * @param user 
 * @returns 
 */
async function addOne(user: IUser): Promise<number | undefined> {
    const resp = await knex('users').insert<IUser>(user).returning<IUser[]>("*");
    return resp[0].id;
}


/**
 * Add user credentials.
 * 
 * @param creds 
 */
function addCreds(creds: IUserCreds): Promise<void> {
    return knex<IUserCreds>('userCreds').insert(creds);
}


/**
 * Fetch all users.
 * 
 * @returns
 */
function fetchAll(): Promise<IUser[]> {
    return knex<IUser, IUser[]>('users');
}


// Export default
export default {
    findByEmail,
    fetchPwdHash,
    addOne,
    addCreds,
    fetchAll,
} as const;
