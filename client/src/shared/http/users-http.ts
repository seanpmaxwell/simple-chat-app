import http from './http';


/**** Vars/Consts ****/

// User routes
const prefix = '/api/users',
    routes = {
        add: ['POST', prefix],
        fetchAll: ['GET', prefix],
    } as const;


// **** Types **** //

export interface IUser {
    id: number;
    email: string;
    name: string;
}


// **** Functions **** //

/**
 * Login user. Return a boolean letting user know if passed or failed.
 */
function add(email: string, name: string, password: string): Promise<void> {
    return http(routes.add, {email, name, password});
}

/**
 * Fetch all users.
 */
function fetchAll(): Promise<{users: IUser[]}> {
    return http(routes.fetchAll);
}


// Export default
export default {
    add,
    fetchAll,
} as const;
