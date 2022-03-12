import { createSlice } from '@reduxjs/toolkit';


// **** Types **** //

export interface IReduxSliceState {
    count: number;
}


// **** Functions **** //

// createSlize();
const slice = createSlice({
    name: 'redux',
    initialState: {
        count: 0,
    } as IReduxSliceState,
    reducers: {
        increment: state => {
            state.count = state.count + 1;
        },
        decrement: state => {
            state.count = state.count - 1;
        },
        reset: state => {
            state.count = 0;
        },
    },
});


// TODO
// **** When you want to starting defining functions outside the object above **** //
// type State = number
// const increment: CaseReducer<State, PayloadAction<number>> = (state, action) =>
//   state + action.payload

// createSlice({
//   name: 'test',
//   initialState: 0,
//   reducers: {
//     increment,
//   },
// })


// Export reducer and actions
export const actions = slice.actions;
export default slice.reducer;
