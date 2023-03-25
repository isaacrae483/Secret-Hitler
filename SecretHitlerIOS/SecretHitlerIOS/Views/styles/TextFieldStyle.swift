//
// Created by Isaac Rae on 3/25/23.
//

import Foundation
import SwiftUI

struct RoundedTextFieldStyle: TextFieldStyle {
    func _body(configuration: TextField<Self._Label>) -> some View {
        configuration
                .padding()
                .frame(maxWidth: .infinity)
                .overlay(
                        RoundedRectangle(cornerRadius: 10)
                                .stroke(Color.gray, lineWidth: 1)
                )
    }
}

extension TextFieldStyle where Self == RoundedTextFieldStyle {
    static var rounded: Self {
        return .init()
    }
}