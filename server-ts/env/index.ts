import path from 'path';
import dotenv from 'dotenv';
import logger from 'jet-logger';
import envVars from '@shared/env-vars';


// Import environment variabes
(() => {
    try {
        const envFolderPath = path.join(__dirname, `${envVars.nodeEnv}.env`);
        const result2 = dotenv.config({
            path: (envFolderPath),
        });
        // Throw error if one
        if (result2.error) {
            throw result2.error;
        }
    } catch (err) {
        logger.err(err);
        process.exit();
    }
})();

