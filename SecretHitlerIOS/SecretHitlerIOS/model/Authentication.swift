//
// Created by Isaac Rae on 3/22/23.
//

import Foundation

struct Authentication: Codable {
    var username: String
    var password: String

    enum CodingKeys: String, CodingKey {
        case username = "username"
        case password = "password"
    }
}
