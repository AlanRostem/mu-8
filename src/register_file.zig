const RegisterFile = struct {
    general_purpose_registers: u8[0x10],
    index_register: u16,
    program_counter: u16,
    stack_pointer: u16,
    delay_timer_register: u8,
    sound_timer_register: u8,
    current_operand: u16,
};
