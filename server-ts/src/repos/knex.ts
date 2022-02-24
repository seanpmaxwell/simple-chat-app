
/**
 * Initialize knex.
 * 
 * created by Sean Maxwell, 1/29/2022
 */

import Knex from 'knex';
import knexConfig from '../../knexfile';


// Initialize knex.
const knex = Knex(knexConfig.development);
export default knex;

