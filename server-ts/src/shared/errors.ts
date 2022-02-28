import HttpStatusCodes from 'http-status-codes';


export abstract class CustomError extends Error {

    public readonly status = HttpStatusCodes.BAD_REQUEST;

    public constructor(msg: string, httpStatus: number) {
        super(msg);
        this.status = httpStatus;
    }
}


export class ParamMissingError extends CustomError {

    public static readonly Msg = 'One or more of the required parameters was missing.';
    public static readonly HttpStatus = HttpStatusCodes.BAD_REQUEST;

    public constructor() {
        super(ParamMissingError.Msg, ParamMissingError.HttpStatus);
    }
}
