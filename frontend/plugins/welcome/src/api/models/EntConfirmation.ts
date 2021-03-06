/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntConfirmationEdges,
    EntConfirmationEdgesFromJSON,
    EntConfirmationEdgesFromJSONTyped,
    EntConfirmationEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntConfirmation
 */
export interface EntConfirmation {
    /**
     * Bookingdate holds the value of the "bookingdate" field.
     * @type {string}
     * @memberof EntConfirmation
     */
    bookingdate?: string;
    /**
     * Bookingend holds the value of the "bookingend" field.
     * @type {string}
     * @memberof EntConfirmation
     */
    bookingend?: string;
    /**
     * Bookingstart holds the value of the "bookingstart" field.
     * @type {string}
     * @memberof EntConfirmation
     */
    bookingstart?: string;
    /**
     * 
     * @type {EntConfirmationEdges}
     * @memberof EntConfirmation
     */
    edges?: EntConfirmationEdges;
    /**
     * Hourstime holds the value of the "hourstime" field.
     * @type {number}
     * @memberof EntConfirmation
     */
    hourstime?: number;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntConfirmation
     */
    id?: number;
}

export function EntConfirmationFromJSON(json: any): EntConfirmation {
    return EntConfirmationFromJSONTyped(json, false);
}

export function EntConfirmationFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntConfirmation {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'bookingdate': !exists(json, 'bookingdate') ? undefined : json['bookingdate'],
        'bookingend': !exists(json, 'bookingend') ? undefined : json['bookingend'],
        'bookingstart': !exists(json, 'bookingstart') ? undefined : json['bookingstart'],
        'edges': !exists(json, 'edges') ? undefined : EntConfirmationEdgesFromJSON(json['edges']),
        'hourstime': !exists(json, 'hourstime') ? undefined : json['hourstime'],
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntConfirmationToJSON(value?: EntConfirmation | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'bookingdate': value.bookingdate,
        'bookingend': value.bookingend,
        'bookingstart': value.bookingstart,
        'edges': EntConfirmationEdgesToJSON(value.edges),
        'hourstime': value.hourstime,
        'id': value.id,
    };
}


