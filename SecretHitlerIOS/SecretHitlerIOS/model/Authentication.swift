//
// Created by Isaac Rae on 3/22/23.
//

import Foundation

struct Authentication: Codable {
    let username: String
    let password: String

    enum CodingKeys: String, CodingKey {
        case username = "username"
        case password = "password"
    }
}
