import { useEffect  } from 'react';
import { Outlet, useNavigate, useLocation } from 'react-router-dom';

import { styled } from '@mui/material/styles';
import Menu from '@mui/material/Menu';
import MenuItem from '@mui/material/MenuItem';

import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Tabs from '@mui/material/Tabs';
import Tab from '@mui/material/Tab';
import { useSetState } from '../shared/hooks';
import authHttp, { ISessionData } from '../shared/http/auth-http';
import { getInitials } from '../shared/functions';


// Constants
const routes = ['/users', '/chat'];



interface IProps {
    sessionData: ISessionData;
    fetchSessionData: () => Promise<void>;
}

interface IState {
    selectedTab: number;
    menuAnchor: HTMLElement | null;
}

function TopBar(props: IProps) {
    const [state, setState ] = useSetState(init()),
        location = useLocation(),
        navigate = useNavigate();
    // On load
    useEffect(() => {
        const loggedIn = (props.sessionData.id > -1),
            path = location.pathname;
        // Navigate to home if not logged in and not at home.
        if (!props.sessionData.waiting && !loggedIn && path !== '/') {
            navigate('/');
        }
        // Listen for changes in the route.
        let idxOfRouteInUrl = routes.indexOf(location.pathname);
        if (idxOfRouteInUrl === -1 || !loggedIn) {
            idxOfRouteInUrl = 0;
        }
        setState({selectedTab: idxOfRouteInUrl});
    }, [location, navigate, props.sessionData.id, props.sessionData.waiting, setState]);
    // Return
    return (
        <>
            <Box sx={{ bgcolor: 'background.paper'}}>
                <AppBar position="static">
                    <Tabs
                        aria-label="simple tabs example"
                        value={state.selectedTab}
                        indicatorColor="secondary"
                        textColor="inherit"
                        onChange={(_, newVal) => {
                            if (newVal === -1) {
                                return;
                            }
                            setState({selectedTab: newVal});
                            navigate(routes[newVal]);
                        }}
                    >
                        {props.sessionData.id !== -1 ? [
                            <NavTab
                                value={0}
                                key={0}
                                label='Users'
                                {...a11yProps(0)}
                            />,
                            <NavTab
                                value={1}
                                key={1}
                                label='Chat'
                                {...a11yProps(1)}
                            />,
                            <NavTab
                                value={-1}
                                key={2} 
                                label={getInitials(props.sessionData.name)}
                                {...a11yProps(2)}
                                sx={{
                                    position: 'absolute',
                                    right: '0'
                                }}
                                onClick={e => {
                                    e.preventDefault();
                                    setState({menuAnchor: e.currentTarget});
                                }}
                            />
                        ] : [
                            <NavTab
                                key={0}
                                label='Home'
                                {...a11yProps(0)}
                            />
                        ]}
                        
                    </Tabs>
                </AppBar>
                <Menu
                    id="basic-menu"
                    anchorEl={state.menuAnchor}
                    open={!!state.menuAnchor}
                    onClose={() => setState({menuAnchor: null})}
                    MenuListProps={{
                    'aria-labelledby': 'basic-button',
                    }}
                >
                    <MenuItem onClick={async () => {
                        const done = await logout();
                        if (done) {
                            setState({menuAnchor: null});
                            await props.fetchSessionData();
                            navigate('/');
                        }
                    }}>
                        Logout
                    </MenuItem>
                </Menu>
            </Box>
            <Outlet />
      </>
    );
}

const NavTab = styled(Tab)({
    fontWeight: 'bold',
});


function a11yProps(index: any) {
    return {
      id: `simple-tab-${index}`,
      'aria-controls': `simple-tabpanel-${index}`,
    };
}


// Get Initial State
function init(): IState {
    return {
        selectedTab: 0,
        menuAnchor: null,
    };
}

// Logout a user
async function logout(): Promise<boolean> {
    let done = false;
    try {
        await authHttp.logout();
        done = true;
    } catch (err) {
        console.error(err);
    }
    return done;
}


// Export default
export default TopBar;
