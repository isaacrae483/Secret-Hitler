//
// Created by Isaac Rae on 3/22/23.
//

import Foundation

struct Room: Codable {
    let id: Int
    let createdAt: Date
    let code: String
    let size: Int

    enum CodingKeys: String, CodingKey {
        case id = "ID"
        case createdAt = "created_at"
        case code = "code"
        case size = "size"
    }
}
