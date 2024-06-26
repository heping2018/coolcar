import * as $protobuf from "protobufjs";
/** Namespace coolcar. */
export namespace coolcar {

    /** TripStatus enum. */
    enum TripStatus {
        TS_NOT_SPECIFIED = 0,
        NOT_STARTED = 1,
        IN_PROGRESS = 2,
        FINISHED = 3,
        PAID = 4
    }

    /** Properties of a Location. */
    interface ILocation {

        /** Location latitude */
        latitude?: (number|null);

        /** Location longitude */
        longitude?: (number|null);
    }

    /** Represents a Location. */
    class Location implements ILocation {

        /**
         * Constructs a new Location.
         * @param [properties] Properties to set
         */
        constructor(properties?: coolcar.ILocation);

        /** Location latitude. */
        public latitude: number;

        /** Location longitude. */
        public longitude: number;

        /**
         * Encodes the specified Location message. Does not implicitly {@link coolcar.Location.verify|verify} messages.
         * @param message Location message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: coolcar.ILocation, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Creates a Location message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns Location
         */
        public static fromObject(object: { [k: string]: any }): coolcar.Location;

        /**
         * Creates a plain object from a Location message. Also converts values to other types if specified.
         * @param message Location
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: coolcar.Location, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this Location to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }

    /** Properties of a Trip. */
    interface ITrip {

        /** Trip start */
        start?: (string|null);

        /** Trip startPos */
        startPos?: (coolcar.ILocation|null);

        /** Trip endPos */
        endPos?: (coolcar.ILocation|null);

        /** Trip pathLocations */
        pathLocations?: (coolcar.ILocation[]|null);

        /** Trip end */
        end?: (string|null);

        /** Trip durationSec */
        durationSec?: (number|null);

        /** Trip feeCent */
        feeCent?: (number|null);

        /** Trip status */
        status?: (coolcar.TripStatus|null);
    }

    /** Represents a Trip. */
    class Trip implements ITrip {

        /**
         * Constructs a new Trip.
         * @param [properties] Properties to set
         */
        constructor(properties?: coolcar.ITrip);

        /** Trip start. */
        public start: string;

        /** Trip startPos. */
        public startPos?: (coolcar.ILocation|null);

        /** Trip endPos. */
        public endPos?: (coolcar.ILocation|null);

        /** Trip pathLocations. */
        public pathLocations: coolcar.ILocation[];

        /** Trip end. */
        public end: string;

        /** Trip durationSec. */
        public durationSec: number;

        /** Trip feeCent. */
        public feeCent: number;

        /** Trip status. */
        public status: coolcar.TripStatus;

        /**
         * Encodes the specified Trip message. Does not implicitly {@link coolcar.Trip.verify|verify} messages.
         * @param message Trip message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: coolcar.ITrip, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Creates a Trip message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns Trip
         */
        public static fromObject(object: { [k: string]: any }): coolcar.Trip;

        /**
         * Creates a plain object from a Trip message. Also converts values to other types if specified.
         * @param message Trip
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: coolcar.Trip, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this Trip to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }

    /** Properties of a GetTripRequest. */
    interface IGetTripRequest {

        /** GetTripRequest id */
        id?: (string|null);
    }

    /** Represents a GetTripRequest. */
    class GetTripRequest implements IGetTripRequest {

        /**
         * Constructs a new GetTripRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: coolcar.IGetTripRequest);

        /** GetTripRequest id. */
        public id: string;

        /**
         * Encodes the specified GetTripRequest message. Does not implicitly {@link coolcar.GetTripRequest.verify|verify} messages.
         * @param message GetTripRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: coolcar.IGetTripRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Creates a GetTripRequest message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns GetTripRequest
         */
        public static fromObject(object: { [k: string]: any }): coolcar.GetTripRequest;

        /**
         * Creates a plain object from a GetTripRequest message. Also converts values to other types if specified.
         * @param message GetTripRequest
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: coolcar.GetTripRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this GetTripRequest to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }

    /** Properties of a GettripResponse. */
    interface IGettripResponse {

        /** GettripResponse id */
        id?: (string|null);

        /** GettripResponse trip */
        trip?: (coolcar.ITrip|null);
    }

    /** Represents a GettripResponse. */
    class GettripResponse implements IGettripResponse {

        /**
         * Constructs a new GettripResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: coolcar.IGettripResponse);

        /** GettripResponse id. */
        public id: string;

        /** GettripResponse trip. */
        public trip?: (coolcar.ITrip|null);

        /**
         * Encodes the specified GettripResponse message. Does not implicitly {@link coolcar.GettripResponse.verify|verify} messages.
         * @param message GettripResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: coolcar.IGettripResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Creates a GettripResponse message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns GettripResponse
         */
        public static fromObject(object: { [k: string]: any }): coolcar.GettripResponse;

        /**
         * Creates a plain object from a GettripResponse message. Also converts values to other types if specified.
         * @param message GettripResponse
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: coolcar.GettripResponse, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this GettripResponse to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }

    /** Represents a TripService */
    class TripService extends $protobuf.rpc.Service {

        /**
         * Constructs a new TripService service.
         * @param rpcImpl RPC implementation
         * @param [requestDelimited=false] Whether requests are length-delimited
         * @param [responseDelimited=false] Whether responses are length-delimited
         */
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);

        /**
         * Calls GetTrip.
         * @param request GetTripRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and GettripResponse
         */
        public getTrip(request: coolcar.IGetTripRequest, callback: coolcar.TripService.GetTripCallback): void;

        /**
         * Calls GetTrip.
         * @param request GetTripRequest message or plain object
         * @returns Promise
         */
        public getTrip(request: coolcar.IGetTripRequest): Promise<coolcar.GettripResponse>;
    }

    namespace TripService {

        /**
         * Callback as used by {@link coolcar.TripService#getTrip}.
         * @param error Error, if any
         * @param [response] GettripResponse
         */
        type GetTripCallback = (error: (Error|null), response?: coolcar.GettripResponse) => void;
    }
}
