import http from './http';


/**** Vars/Consts ****/

// Routes
const prefix = '/api/auth',
    routes = {
        login: ['PUT', `${prefix}/login`],
        logout: ['GET', `${prefix}/logout`],
        sessionData: ['GET', `${prefix}/session-data`],
    } as const;


/**** Types ****/

// What's stored in a login session
export interface ISessionData {
    id: number;
    email: string;
    name: string;
    waiting: boolean;
}


/**** Functions ****/

/**
 * Login user. Return a boolean letting user know if passed or failed.
 */
async function login(email: string, password: string): Promise<boolean> {
    const resp = await http(routes.login, {email, password});
    return resp.passed;
}

/**
 * Delete the session cookie.
 */
function logout(): Promise<void> {
    return http(routes.logout);
}

/**
 * Fetch session data from jwt.
 */
function getSessionData(): Promise<ISessionData> {
    return http(routes.sessionData);
}   


// Export default
export default {
    login,
    logout,
    getSessionData,
} as const;
