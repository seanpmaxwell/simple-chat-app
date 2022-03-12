import axios, { Method } from 'axios';


/**** Functions ****/

/**
 * Wrapper for the axios function.
 */
async function http(route: Readonly<string[]>, data?: Record<string, any>) {
    const resp = await axios({method: route[0] as Method, url: route[1], data});
    return resp.data;
}


// Export default
export default http;
