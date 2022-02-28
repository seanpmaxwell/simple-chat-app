// User Interface
export interface IUser {
    id?: number;
    email: string;
    name: string;
    createdAt: string;
}

// User credentials to store the password hash.
export interface IUserCreds {
    id?: number;
    pwdHash: string;
    userId: number;
}

// Users saved in db will not have undefined id
export type TSavedUser = Required<IUser>;


/**
 * Get new user instance.
 */
function getNew(
    email?: string,
    name?: string,
    createdAt?: string,
): IUser {
    return {
        email: email ?? '',
        name: name ?? '',
        createdAt: createdAt ?? new Date().toISOString(),
    };
}


/**
 * Copy a user object.
 */
function copy(user: IUser): IUser {
    return {
        id: user?.id ?? -1,
        email: user?.email ?? '',
        name: user?.name ?? '',
        createdAt: user?.createdAt ?? new Date().toISOString(),
    };
}


// Export user functions
export default {
    new: getNew,
    copy,
} as const;
