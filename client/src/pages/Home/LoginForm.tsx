import { useContext } from 'react';
import { useNavigate } from 'react-router-dom';

import {
    Button,
    Grid,
    TextField,
    Typography,
} from '@mui/material';

import { useSetState } from '../../shared/hooks';
import authHttp from '../../shared/http/auth-http';
import { appCtx } from '../App';


// **** Types **** //

// Login Status
enum LoginStatus {
    NA = 1,
    Waiting,
    Failed,
    Passed,
}

// LoginForm State
interface IState {
    email: string;
    password: string;
    loginStatus: LoginStatus
}


// **** LoginForm **** //

/**
 * Render()
 */
function LoginForm(): JSX.Element {
    const [state, setState ] = useSetState(init()),
    ctx = useContext(appCtx),
    navigate = useNavigate();
    // Return
    return (
        <Grid
            container={true}
            direction="column"
            justifyContent="center"
            alignItems="center"
            spacing={3}
        >
            <Grid item={true}>
                <Typography
                    variant="h5"
                    gutterBottom={true}
                >
                    Login
                </Typography>
            </Grid>
            <Grid item={true}>
                <TextField
                    id="login-email"
                    label="Email"
                    type="text"
                    value={state.email}
                    onChange={(e) => setState({email: e.currentTarget.value})}
                />
            </Grid>
            <Grid item={true}>
                <TextField
                    id="login-password"
                    label="Password"
                    type="password"
                    value={state.password}
                    onChange={(e) => setState({password: e.currentTarget.value})}
                />
            </Grid>
            <br/>
            <Grid item={true}>
                <Button
                    color='primary'
                    variant='contained'
                    disabled={!state.email || !state.password}
                    onClick={async () => {
                        setState({loginStatus: LoginStatus.Waiting});
                        const loginStatus = await loginUser(state);
                        if (loginStatus === LoginStatus.Passed) {
                            await ctx.fetchSessionData();
                            navigate('/users');
                        } else {
                            setState({loginStatus});
                        }
                    }}
                >
                    Login
                </Button>
            </Grid>
        </Grid>
    );
}


/**** Functions ****/

/**
 * Initial State
 */
function init(): IState {
    return {
        email: 'foo@barr',
        password: 'Password@1',
        loginStatus: LoginStatus.NA,
    };
}

/**
 * Login User
 */
async function loginUser(state: IState): Promise<LoginStatus> {
    const { email, password } = state;
    let status = LoginStatus.Failed;
    try {
        const passed = await authHttp.login(email, password);
        status = (passed ? LoginStatus.Passed : LoginStatus.Failed);
    } catch (err) {
        console.error(err);
    }
    return status;
}


// Export default
export default LoginForm;
