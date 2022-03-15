import React, { useCallback, useEffect } from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { configureStore } from '@reduxjs/toolkit';
import { Provider} from 'react-redux';

import CssBaseline from '@mui/material/CssBaseline';

import Home from './Home/Home';
import NoPage from './NoPage/NoPage';
import './App.css';
import TopBar from './TopBar/Topbar';
import Users from './Users/Users';
import Chat from './Chat/Chat';
import Redux from './Redux/Redux';
import authHttp, { ISessionData } from '../shared/http/auth-http';
import { useSetState } from '../shared/hooks';
import ReduxSlice from './Redux/ReduxSlice';


// **** Setup Redux **** //

const store = configureStore({
    reducer: ReduxSlice,
});


// **** Setup Context **** //

export let appCtx = React.createContext<IContext>({
    sessionData: getEmptySessionData(),
    fetchSessionData: () => Promise.resolve(),
});

interface IContext {
    sessionData: ISessionData
    fetchSessionData: () => Promise<void>;
}


// **** Types **** //

interface IState {
    sessionData: ISessionData;
}


// **** App **** //

/**
 * Render()
 */
function App() {
    const [state, setState] = useSetState(init());
    // Set fetch-session-data function
    const fetchSessionData = useCallback(() => {
        return getSessionData().then(sd => setState({sessionData: sd}));
    }, [setState]);
    // On load
    useEffect(() => {
        fetchSessionData();
    }, [fetchSessionData]);
    // Return
    return (
        <div>
            <Provider store={store}>
                <appCtx.Provider
                    value={{
                        sessionData: state.sessionData,
                        fetchSessionData,
                    }}
                >
                    <React.Fragment>
                        <CssBaseline />
                        <BrowserRouter>
                            <Routes>
                                <Route
                                    path="/"
                                    element={
                                        <TopBar/>
                                    }
                                >
                                    <Route
                                        index={true}
                                        element={
                                            <Home/>
                                        }
                                    />
                                    <Route path="redux" element={<Redux />} />
                                    <Route path="users" element={<Users />} />
                                    <Route path="chat" element={<Chat />} />
                                    <Route path="*" element={<NoPage />} />
                                </Route>
                            </Routes>
                        </BrowserRouter>
                    </React.Fragment>
                </appCtx.Provider>
            </Provider>
        </div>
    );
}


/**** Functions ****/

/**
 * Get initial-state
 */
function init(): IState {
    return {
        sessionData: getEmptySessionData(),
    };
}

/**
 * Get blank session-data (avoid undefined errors before login)
 */
function getEmptySessionData(): ISessionData {
    return {
        id: -1,
        email: '',
        name: '',
        waiting: true,
    };
}

/**
 * Fetch jwt session-data
 */
async function getSessionData(): Promise<ISessionData> {
    try {
        const data = await authHttp.getSessionData();
        return data;
    } catch (err) {
        console.error(err);
    }
    return getEmptySessionData();
}


// Export default
export default App;
